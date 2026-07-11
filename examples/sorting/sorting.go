// تنفذ الحزمة `slices` في Go فرز الأنواع المدمجة والأنواع
// المعرّفة من المستخدم. سنتعرف أولًا على فرز الأنواع المدمجة.

package main

import (
	"fmt"
	"slices"
)

func main() {

	// دوال الفرز عامة، وتعمل مع أي نوع مدمج _قابل للترتيب_. راجع
	// [cmp.Ordered](https://pkg.go.dev/cmp#Ordered) للاطلاع على
	// قائمة الأنواع القابلة للترتيب.
	strs := []string{"c", "a", "b"}
	slices.Sort(strs)
	fmt.Println("Strings:", strs)

	// هذا مثال على فرز قيم من النوع `int`.
	ints := []int{7, 2, 4}
	slices.Sort(ints)
	fmt.Println("Ints:   ", ints)

	// يمكننا أيضًا استخدام الحزمة `slices` للتحقق مما إذا كانت
	// شريحة مرتبة بالفعل.
	s := slices.IsSorted(ints)
	fmt.Println("Sorted: ", s)
}
