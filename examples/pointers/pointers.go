// تدعم Go <em><a href="https://en.wikipedia.org/wiki/Pointer_(computer_programming)">المؤشرات</a></em>،
// مما يتيح لك تمرير مراجع إلى القيم والسجلات داخل
// برنامجك.

package main

import "fmt"

// سنوضح الفرق بين عمل المؤشرات والقيم باستخدام دالتين:
// `zeroval` و`zeroptr`. تستقبل `zeroval` معاملًا من النوع
// `int`، لذلك ستمرر الوسائط إليها بالقيمة. وستحصل
// `zeroval` على نسخة من `ival` تختلف عن النسخة الموجودة
// في الدالة المستدعية.
func zeroval(ival int) {
	ival = 0
}

// في المقابل، تستقبل `zeroptr` معاملًا من النوع `*int`،
// ما يعني أنها تستقبل مؤشرًا إلى `int`. تفك الصيغة
// `*iptr` في جسم الدالة إشارة المؤشر، فتصل من عنوان
// الذاكرة إلى القيمة الحالية المخزنة فيه. يؤدي إسناد قيمة
// إلى مؤشر مفكوك الإشارة إلى تغيير القيمة في العنوان
// المشار إليه.
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	// تعطي الصيغة `&i` عنوان ذاكرة `i`، أي مؤشرًا إلى `i`.
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	// يمكن طباعة المؤشرات أيضًا.
	fmt.Println("pointer:", &i)

	// يمكن إنشاء مؤشر جديد إلى قيمة باستخدام الدالة
	// المدمجة `new`.
	p := new(42)
	fmt.Println("value at *p:", *p)
	zeroptr(p)
	fmt.Println("value at *p:", *p)
}
