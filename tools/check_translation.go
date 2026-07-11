// Checks Arabic translation metadata and source coverage.
package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

var (
	hyphenPattern   = regexp.MustCompile(`-+`)
	progressPattern = regexp.MustCompile("^\\|\\s*(\\d+)\\s*\\|\\s*`([^`]+)`\\s*\\|[^|]+\\|\\s*(✅|⬜)\\s*\\|\\s*(✅|⬜)\\s*\\|\\s*$")
	technicalTitles = map[string]bool{"JSON": true, "XML": true}
)

type translationChecker struct {
	failures []string
}

func (c *translationChecker) failf(format string, args ...any) {
	c.failures = append(c.failures, fmt.Sprintf(format, args...))
}

func readLines(path string, c *translationChecker) []string {
	file, err := os.Open(path)
	if err != nil {
		c.failf("%s: %v", path, err)
		return nil
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		c.failf("%s: %v", path, err)
	}
	return lines
}

func exampleIDFromName(name string) string {
	id := strings.ToLower(strings.TrimSpace(name))
	id = strings.ReplaceAll(id, " ", "-")
	id = strings.ReplaceAll(id, "/", "-")
	id = strings.ReplaceAll(id, "'", "")
	return hyphenPattern.ReplaceAllString(id, "-")
}

func containsArabic(text string) bool {
	for _, r := range text {
		if unicode.In(r, unicode.Arabic) {
			return true
		}
	}
	return false
}

func isBidiControl(r rune) bool {
	return r == '\u200e' || r == '\u200f' ||
		(r >= '\u202a' && r <= '\u202e') ||
		(r >= '\u2066' && r <= '\u2069')
}

func checkNoBidiControls(path string, data []byte, c *translationChecker) {
	line := 1
	for _, r := range string(data) {
		if isBidiControl(r) {
			c.failf("%s:%d: hidden bidi control U+%04X is not allowed", path, line, r)
		}
		if r == '\n' {
			line++
		}
	}
}

func readExpectedExampleIDs(c *translationChecker) []string {
	var ids []string
	seen := make(map[string]bool)
	for lineNumber, rawLine := range readLines("examples.txt", c) {
		line := strings.TrimSpace(rawLine)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		id := exampleIDFromName(line)
		if seen[id] {
			c.failf("examples.txt:%d: duplicate example ID %q", lineNumber+1, id)
			continue
		}
		seen[id] = true
		ids = append(ids, id)
	}
	return ids
}

func checkArabicTitleCatalog(expectedIDs []string, c *translationChecker) {
	expected := make(map[string]bool, len(expectedIDs))
	for _, id := range expectedIDs {
		expected[id] = true
	}

	var orderedIDs []string
	seen := make(map[string]bool)
	for lineNumber, rawLine := range readLines("examples.ar.txt", c) {
		line := strings.TrimSpace(rawLine)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "|", 2)
		if len(parts) != 2 {
			c.failf("examples.ar.txt:%d: expected example-id|Arabic title", lineNumber+1)
			continue
		}
		id := strings.TrimSpace(parts[0])
		title := strings.TrimSpace(parts[1])
		if !expected[id] {
			c.failf("examples.ar.txt:%d: unknown example ID %q", lineNumber+1, id)
		}
		if seen[id] {
			c.failf("examples.ar.txt:%d: duplicate example ID %q", lineNumber+1, id)
			continue
		}
		seen[id] = true
		orderedIDs = append(orderedIDs, id)
		if !containsArabic(title) && !technicalTitles[title] {
			c.failf("examples.ar.txt:%d: title for %q must contain Arabic text or be a technical initialism", lineNumber+1, id)
		}
	}

	for index, id := range expectedIDs {
		if !seen[id] {
			c.failf("examples.ar.txt: missing Arabic title for example ID %q", id)
		}
		if index < len(orderedIDs) && orderedIDs[index] != id {
			c.failf("examples.ar.txt: entry %d is %q; expected %q", index+1, orderedIDs[index], id)
		}
	}
	if len(orderedIDs) != len(expectedIDs) {
		c.failf("examples.ar.txt: found %d entries; expected %d", len(orderedIDs), len(expectedIDs))
	}
}

func checkExampleSources(expectedIDs []string, c *translationChecker) {
	for _, id := range expectedIDs {
		dir := filepath.Join("examples", id)
		hasArabicDocs := false
		hasSource := false
		err := filepath.WalkDir(dir, func(path string, entry os.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if entry.IsDir() || (filepath.Ext(path) != ".go" && filepath.Ext(path) != ".sh") {
				return nil
			}
			hasSource = true
			data, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			for _, line := range strings.Split(string(data), "\n") {
				trimmed := strings.TrimSpace(line)
				if (strings.HasPrefix(trimmed, "//") || strings.HasPrefix(trimmed, "#")) && containsArabic(trimmed) {
					hasArabicDocs = true
				}
			}
			return nil
		})
		if err != nil {
			c.failf("%s: %v", dir, err)
			continue
		}
		if !hasSource {
			c.failf("%s: no Go or shell source found", dir)
		}
		if !hasArabicDocs {
			c.failf("%s: no Arabic explanatory comment found", dir)
		}
	}
}

func checkProgress(expectedIDs []string, c *translationChecker) {
	seen := make(map[string]bool)
	position := 0
	for lineNumber, line := range readLines("PROGRESS.md", c) {
		match := progressPattern.FindStringSubmatch(line)
		if match == nil {
			continue
		}
		number, _ := strconv.Atoi(match[1])
		id := match[2]
		if seen[id] {
			c.failf("PROGRESS.md:%d: duplicate example ID %q", lineNumber+1, id)
			continue
		}
		seen[id] = true
		position++
		if position > len(expectedIDs) || expectedIDs[position-1] != id || number != position {
			c.failf("PROGRESS.md:%d: example order or number does not match examples.txt", lineNumber+1)
		}
		if match[3] != "✅" || match[4] != "✅" {
			c.failf("PROGRESS.md:%d: example %q is not translated and reviewed", lineNumber+1, id)
		}
	}
	for _, id := range expectedIDs {
		if !seen[id] {
			c.failf("PROGRESS.md: missing status row for example ID %q", id)
		}
	}
}

func checkTextFilesForBidiControls(c *translationChecker) {
	textExtensions := map[string]bool{
		".css": true, ".go": true, ".html": true, ".js": true, ".md": true,
		".sh": true, ".tmpl": true, ".txt": true, ".yaml": true, ".yml": true,
	}
	err := filepath.WalkDir(".", func(path string, entry os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if entry.IsDir() && (entry.Name() == ".git" || entry.Name() == ".idea" || entry.Name() == ".vscode") {
			return filepath.SkipDir
		}
		if entry.IsDir() || !textExtensions[filepath.Ext(path)] {
			return nil
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		checkNoBidiControls(path, data, c)
		return nil
	})
	if err != nil {
		c.failf("bidi-control scan: %v", err)
	}
}

func main() {
	checker := &translationChecker{}
	expectedIDs := readExpectedExampleIDs(checker)
	checkArabicTitleCatalog(expectedIDs, checker)
	checkExampleSources(expectedIDs, checker)
	checkProgress(expectedIDs, checker)
	checkTextFilesForBidiControls(checker)

	if len(checker.failures) > 0 {
		for _, failure := range checker.failures {
			fmt.Fprintln(os.Stderr, "translation check:", failure)
		}
		os.Exit(1)
	}
	fmt.Printf("Translation checks passed for %d examples.\n", len(expectedIDs))
}
