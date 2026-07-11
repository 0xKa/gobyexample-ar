// أضافت Go منذ الإصدار 1.18 دعم _الأنواع العامة_، المعروفة
// أيضًا باسم _مُعامِلات الأنواع_.

package main

import "fmt"

// كمثال على دالة عامة، تستقبل `SlicesIndex` شريحة من أي
// نوع `comparable` وعنصرًا من ذلك النوع، ثم تعيد فهرس أول
// ظهور لـ`v` في `s`، أو `-1` إن لم يوجد. يعني القيد
// `comparable` أنه يمكننا مقارنة قيم هذا النوع بالمعاملين
// `==` و`!=`. لشرح أشمل لتوقيع النوع هذا، راجع
// [هذه المقالة](https://go.dev/blog/deconstructing-type-parameters).
// لاحظ أن هذه الدالة موجودة في المكتبة القياسية باسم
// [slices.Index](https://pkg.go.dev/slices#Index).
func SlicesIndex[S ~[]E, E comparable](s S, v E) int {
	for i := range s {
		if v == s[i] {
			return i
		}
	}
	return -1
}

// كمثال على نوع عام، تمثل `List` قائمة أحادية الربط
// بقيم من أي نوع.
type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

// يمكننا تعريف الأساليب على الأنواع العامة كما نفعل مع
// الأنواع العادية، لكن علينا إبقاء مُعامِلات الأنواع في
// مواضعها. النوع هو `List[T]` لا `List`.
func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

// تعيد `AllElements` جميع عناصر `List` في شريحة. سنرى في
// المثال التالي طريقةً أقرب إلى أسلوب Go المتعارف عليه
// لاجتياز جميع عناصر الأنواع المخصصة.
func (lst *List[T]) AllElements() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func main() {
	var s = []string{"foo", "bar", "zoo"}

	// عند استدعاء الدوال العامة، يمكننا غالبًا الاعتماد على
	// _استنتاج النوع_. لاحظ أننا لا نحتاج إلى تحديد نوعي
	// `S` و`E` عند استدعاء `SlicesIndex`، إذ يستنتجهما
	// المترجم تلقائيًا.
	fmt.Println("index of zoo:", SlicesIndex(s, "zoo"))

	// ... مع أنه يمكننا أيضًا تحديدهما صراحةً.
	_ = SlicesIndex[[]string, string](s, "zoo")

	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("list:", lst.AllElements())
}
