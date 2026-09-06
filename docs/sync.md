# Upstream Synchronization Guide

This guide describes the end-to-end workflow for keeping the Arabic translation repository ([`0xKa/gobyexample-ar`](https://github.com/0xKa/gobyexample-ar)) synchronized with the original upstream project ([`mmcgrana/gobyexample`](https://github.com/mmcgrana/gobyexample)).

---

## 1. Terminology & Remotes

* **`origin`**: Your repository (the Arabic translation fork):
  `https://github.com/0xKa/gobyexample-ar.git`
* **`upstream`**: The original Go by Example repository:
  `https://github.com/mmcgrana/gobyexample.git`

Verify your remotes:

```bash
git remote -v
```

If `upstream` is not listed, add it:

```bash
git remote add upstream https://github.com/mmcgrana/gobyexample.git
```

---

## 2. Check for New Upstream Commits

Fetch the latest commits from upstream:

```bash
git fetch upstream
```

Compare your local `master` branch with `upstream/master`:

```bash
git log --oneline HEAD..upstream/master
```

* If the command produces no output, your repository is already up to date.
* If commits are listed, inspect the changes in detail:

```bash
git log -p HEAD..upstream/master
```

Note which examples or configuration files (`go.mod`, etc.) were added or modified.

---

## 3. Create a Sync Branch

Always perform merges in a separate branch so you can resolve conflicts, update translations, and test cleanly:

```bash
git checkout master
git pull origin master
git checkout -b sync/upstream-$(date +%Y-%m-%d)
```

---

## 4. Merge Upstream

Start the merge without automatically committing:

```bash
git merge upstream/master --no-commit
```

Check the status to see which files merged cleanly and which have conflicts:

```bash
git status
```

---

## 5. Resolve Conflicts and Update Arabic Content

### Common Conflict Scenarios:

1. **`README.md`**:
   - Upstream usually updates its English README (e.g., adding links to new language translations under `### Translations`).
   - Our repository maintains a localized Arabic `README.md`.
   - Typically, keep our Arabic version unless there is a project-level change that should be adapted into Arabic:
     ```bash
     git checkout --ours README.md
     ```

2. **Go source files (`examples/*/*.go`)**:
   - Accept the updated Go code and logic from upstream.
   - Update the accompanying Arabic prose comments (`// ...`) to reflect any new or modified explanations.
   - Refer to [`GLOSSARY.md`](GLOSSARY.md) for approved technical terms.
   - Ensure no invisible Unicode bidirectional control characters are introduced.

3. **Shell scripts (`examples/*/*.sh`)**:
   - Accept upstream terminal commands and expected output updates.
   - Update any descriptive `# ...` comments in Arabic if they were changed upstream (e.g., updated URLs).

4. **Go Playground hashes (`examples/*/*.hash`) & generated HTML (`public/*`)**:
   - **Do not edit `public/` files manually.**
   - You can resolve conflicts in `.hash` and `public/` files by taking either version (e.g. `git checkout --theirs examples/.../*.hash public/...`). They will be completely regenerated and recalculated by `tools/build` in the next step.

---

## 6. Build, Test, and Regenerate

### Step A: Run translation and syntax tests

```bash
tools/test
```

This runs `go vet` on all examples and `tools/check_translation.go` to verify translation completeness, catalog alignment, and absence of bidirectional control characters.

### Step B: Regenerate `public/` and refresh Go Playground hashes

Running `tools/build` without `TESTING=1` computes the code hashes of modified Go examples and queries `go.dev/_/share` to get updated playground URLs:

```bash
VERBOSE=1 tools/build
```

### Step C: Verify clean build in testing mode

Run `tools/build` in strict testing mode to guarantee that `public/` matches the source code exactly:

```bash
VERBOSE=1 TESTING=1 tools/build
git diff --exit-code
```

Both commands should succeed with exit code `0`.

---

## 7. Manual Verification

Before pushing, preview the generated website locally:

```bash
tools/serve
```

1. Open `http://127.0.0.1:8000/` in your browser.
2. Navigate to the examples you updated (e.g., `http://127.0.0.1:8000/<example-id>`).
3. Check:
   - Arabic text reads cleanly right-to-left (RTL).
   - Code blocks and terminal outputs remain left-to-right (LTR).
   - The **"Run code"** (Go Playground) link and copy buttons work properly.

Press `Ctrl+C` in your terminal to stop the server when done.

---

## 8. Commit and Push

Stage all resolved files, updated hashes, and regenerated `public/` assets:

```bash
git add -A
git commit -m "Merge upstream master into Arabic translation"
```

Push the sync branch to your repository:

```bash
git push -u origin sync/upstream-$(date +%Y-%m-%d)
```

---

## 9. Merge into `master`

You can either:

* **Option A (Recommended for review):** Open a Pull Request on GitHub from the sync branch into `master` and merge it after review.
* **Option B (Direct fast-forward merge):**
  ```bash
  git checkout master
  git merge --ff-only sync/upstream-$(date +%Y-%m-%d)
  git push origin master
  ```

Once merged into `master`, clean up the sync branch:

```bash
git branch -d sync/upstream-$(date +%Y-%m-%d)
git push origin --delete sync/upstream-$(date +%Y-%m-%d)
```

GitHub Actions will automatically run the CI checks and deploy the updated site to GitHub Pages.
