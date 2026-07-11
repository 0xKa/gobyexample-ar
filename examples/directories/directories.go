// توفر Go عدة دوال مفيدة للعمل مع *المجلدات* في نظام الملفات.

package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// أنشئ مجلدًا فرعيًا جديدًا داخل مجلد العمل الحالي.
	err := os.Mkdir("subdir", 0755)
	check(err)

	// عند إنشاء مجلدات مؤقتة، يُستحسن تأجيل حذفها باستخدام
	// `defer`. تحذف `os.RemoveAll` شجرة مجلدات كاملة، على نحو شبيه
	// بـ`rm -rf`.
	defer os.RemoveAll("subdir")

	// دالة مساعدة لإنشاء ملف فارغ جديد.
	createEmptyFile := func(name string) {
		d := []byte("")
		check(os.WriteFile(name, d, 0644))
	}

	createEmptyFile("subdir/file1")

	// يمكننا إنشاء تسلسل هرمي من المجلدات، بما فيها المجلدات
	// الأم، باستخدام `MkdirAll`. يشبه ذلك أمر سطر الأوامر
	// `mkdir -p`.
	err = os.MkdirAll("subdir/parent/child", 0755)
	check(err)

	createEmptyFile("subdir/parent/file2")
	createEmptyFile("subdir/parent/file3")
	createEmptyFile("subdir/parent/child/file4")

	// تسرد `ReadDir` محتويات المجلد، وتعيد شريحة من كائنات
	// `os.DirEntry`.
	c, err := os.ReadDir("subdir/parent")
	check(err)

	fmt.Println("Listing subdir/parent")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	// تتيح لنا `Chdir` تغيير مجلد العمل الحالي على نحو شبيه بـ`cd`.
	err = os.Chdir("subdir/parent/child")
	check(err)

	// سنرى الآن محتويات `subdir/parent/child` عند سرد المجلد
	// *الحالي*.
	c, err = os.ReadDir(".")
	check(err)

	fmt.Println("Listing subdir/parent/child")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	// عُد باستخدام `cd` إلى موضع البداية.
	err = os.Chdir("../../..")
	check(err)

	// يمكننا أيضًا زيارة مجلد *بكل تفرعاته*، بما في ذلك جميع
	// مجلداته الفرعية. تقبل `WalkDir` دالة رد نداء لمعالجة كل ملف
	// أو مجلد تزوره.
	fmt.Println("Visiting subdir")
	err = filepath.WalkDir("subdir", visit)
	check(err)
}

// تُستدعى `visit` لكل ملف أو مجلد تعثر عليه `filepath.WalkDir`
// في جميع التفرعات.
func visit(path string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}
	fmt.Println(" ", path, d.IsDir())
	return nil
}
