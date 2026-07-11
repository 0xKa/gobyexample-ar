// تتبع كتابة الملفات في Go أنماطًا شبيهة بما رأيناه سابقًا عند
// القراءة.

package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// للبدء، هكذا نكتب سلسلة نصية، أو مجرد بايتات، إلى ملف.
	d1 := []byte("hello\ngo\n")
	path1 := filepath.Join(os.TempDir(), "dat1")
	err := os.WriteFile(path1, d1, 0644)
	check(err)

	// لعمليات كتابة أدق، افتح ملفًا للكتابة.
	path2 := filepath.Join(os.TempDir(), "dat2")
	f, err := os.Create(path2)
	check(err)

	// من الشائع تأجيل `Close` مباشرة بعد فتح ملف.
	defer f.Close()

	// يمكنك كتابة شرائح البايتات باستخدام `Write` كما هو متوقع.
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	// تتوفر أيضًا `WriteString`.
	n3, err := f.WriteString("writes\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n3)

	// استدعِ `Sync` لتفريغ عمليات الكتابة إلى التخزين الدائم.
	f.Sync()

	// توفر `bufio` كائنات `Writer` مخزنة مؤقتًا، إضافة إلى كائنات
	// `Reader` المخزنة مؤقتًا التي رأيناها سابقًا.
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n4)

	// استخدم `Flush` لضمان تطبيق جميع العمليات المخزنة مؤقتًا على
	// الكاتب الأساسي.
	w.Flush()

}
