// _الواجهات_ مجموعات مسماة من تواقيع الأساليب.

package main

import (
	"fmt"
	"math"
)

// إليك واجهة أساسية للأشكال الهندسية.
type geometry interface {
	area() float64
	perim() float64
}

// سنطبق هذه الواجهة في مثالنا على النوعين `rect`
// و`circle`.
type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

// لتطبيق واجهة في Go، يكفي تطبيق جميع الأساليب الموجودة
// فيها. نطبق هنا `geometry` على قيم `rect`.
func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// وهذا هو التطبيق لقيم `circle`.
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// إذا كان للمتغير نوع واجهة، فيمكننا استدعاء الأساليب
// الموجودة في الواجهة المسماة. تستفيد الدالة `measure`
// هنا من ذلك للعمل على أي قيمة `geometry`.
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

// يفيد أحيانًا معرفة نوع قيمة الواجهة وقت التشغيل. أحد
// الخيارات هو استخدام *تأكيد النوع* كما هو موضح هنا،
// والآخر استخدام [عبارة `switch` للأنواع](switch).
func detectCircle(g geometry) {
	if c, ok := g.(circle); ok {
		fmt.Println("circle with radius", c.radius)
	}
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	// يطبق نوعا الهيكل `circle` و`rect` واجهة `geometry`،
	// لذا يمكننا تمرير قيم من هذين الهيكلين كوسائط إلى
	// `measure`.
	measure(r)
	measure(c)

	detectCircle(r)
	detectCircle(c)
}
