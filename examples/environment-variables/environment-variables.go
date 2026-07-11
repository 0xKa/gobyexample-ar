// [متغيرات البيئة](https://en.wikipedia.org/wiki/Environment_variable)
// آلية عامة [لنقل معلومات الإعداد إلى برامج
// Unix](https://www.12factor.net/config). لنتعرف على كيفية تعيين
// متغيرات البيئة والحصول عليها وسردها.

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	// لتعيين زوج مفتاح وقيمة، استخدم `os.Setenv`. وللحصول على قيمة
	// مفتاح، استخدم `os.Getenv`. ستعيد `os.Getenv` سلسلة نصية فارغة
	// إذا لم يكن المفتاح موجودًا في البيئة.
	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	// استخدم `os.Environ` لسرد جميع أزواج المفاتيح والقيم في
	// البيئة. تعيد هذه شريحة سلاسل نصية بالصيغة `KEY=value`.
	// يمكنك استخدام `strings.SplitN` لتقسيمها والحصول على المفتاح
	// والقيمة. نطبع هنا جميع المفاتيح.
	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0])
	}
}
