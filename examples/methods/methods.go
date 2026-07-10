// تدعم Go تعريف _الأساليب_ على أنواع الهياكل.

package main

import "fmt"

type rect struct {
	width, height int
}

// للأسلوب `area` _نوع مستقبِل_ هو `*rect`.
func (r *rect) area() int {
	return r.width * r.height
}

// يمكن تعريف الأساليب على أنواع مستقبِلات من المؤشرات أو
// القيم. إليك مثالًا على مستقبِل من نوع قيمة.
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	// نستدعي هنا الأسلوبين المعرّفين لهيكلنا.
	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	// تتولى Go تلقائيًا التحويل بين القيم والمؤشرات عند
	// استدعاء الأساليب. قد ترغب في استخدام نوع مستقبِل من
	// المؤشرات لتجنب النسخ عند استدعاء الأسلوب، أو للسماح
	// للأسلوب بتغيير الهيكل المستقبِل.
	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}
