// قد نرغب أحيانًا في فرز مجموعة وفق معيار غير ترتيبها الطبيعي.
// لنفترض مثلًا أننا نريد فرز السلاسل النصية حسب طولها بدلًا من
// ترتيبها أبجديًا. إليك مثالًا على الفرز المخصص في Go.

package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	fruits := []string{"peach", "banana", "kiwi"}

	// ننفذ دالة مقارنة لأطوال السلاسل النصية. تفيدنا
	// `cmp.Compare` في ذلك.
	lenCmp := func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}

	// يمكننا الآن استدعاء `slices.SortFunc` مع دالة المقارنة
	// المخصصة هذه لفرز `fruits` حسب طول الاسم.
	slices.SortFunc(fruits, lenCmp)
	fmt.Println(fruits)

	// يمكننا استخدام الأسلوب نفسه لفرز شريحة من قيم ليست من
	// الأنواع المدمجة.
	type Person struct {
		name string
		age  int
	}

	people := []Person{
		Person{name: "Jax", age: 37},
		Person{name: "TJ", age: 25},
		Person{name: "Alex", age: 72},
	}

	// افرز `people` حسب العمر باستخدام `slices.SortFunc`.
	//
	// لاحظ أنه إذا كان الهيكل `Person` كبيرًا، فقد تفضل أن تحتوي
	// الشريحة على `*Person` بدلًا منه، وأن تعدّل دالة الفرز وفقًا
	// لذلك. عند الشك، [قِس الأداء](testing-and-benchmarking)!
	slices.SortFunc(people,
		func(a, b Person) int {
			return cmp.Compare(a.age, b.age)
		})
	fmt.Println(people)
}
