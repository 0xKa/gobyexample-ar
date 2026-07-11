// توفر الحزمة `filepath` دوال لتحليل *مسارات الملفات* وإنشائها
// بطريقة قابلة للنقل بين أنظمة التشغيل، مثل `dir/file` على Linux
// مقابل `dir\file` على Windows.
package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {

	// ينبغي استخدام `Join` لإنشاء المسارات بطريقة قابلة للنقل. تقبل
	// أي عدد من الوسائط وتنشئ منها مسارًا هرميًا.
	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println("p:", p)

	// ينبغي دائمًا استخدام `Join` بدلًا من وصل `/` أو `\` يدويًا.
	// إلى جانب قابلية النقل، تطبّع `Join` المسارات أيضًا بإزالة
	// الفواصل الزائدة ومكونات الانتقال بين المجلدات مثل `..`.
	fmt.Println(filepath.Join("dir1//", "filename"))
	fmt.Println(filepath.Join("dir1/../dir1", "filename"))

	// يمكن استخدام `Dir` و`Base` لتقسيم مسار إلى المجلد والملف.
	// أو يمكن استخدام `Split` لإعادتهما معًا في استدعاء واحد.
	fmt.Println("Dir(p):", filepath.Dir(p))
	fmt.Println("Base(p):", filepath.Base(p))

	// يمكننا التحقق مما إذا كان المسار مطلقًا.
	fmt.Println(filepath.IsAbs("dir/file"))
	fmt.Println(filepath.IsAbs("/dir/file"))

	filename := "config.json"

	// لبعض أسماء الملفات امتدادات تأتي بعد نقطة. يمكننا فصل
	// الامتداد عن هذه الأسماء باستخدام `Ext`.
	ext := filepath.Ext(filename)
	fmt.Println(ext)

	// للحصول على اسم الملف بعد إزالة الامتداد، استخدم
	// `strings.TrimSuffix`.
	fmt.Println(strings.TrimSuffix(filename, ext))

	// تجد `Rel` مسارًا نسبيًا بين *أساس* و*هدف*. وتعيد خطأ إذا
	// تعذر جعل الهدف نسبيًا إلى الأساس.
	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
}
