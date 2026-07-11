// أضافت Go منذ الإصدار 1.23 دعم
// [المكررات](https://go.dev/blog/range-functions)،
// مما يتيح لنا استخدام `range` مع أي شيء تقريبًا!

package main

import (
	"fmt"
	"iter"
	"slices"
	"strings"
)

// لننظر مجددًا إلى النوع `List` من
// [المثال السابق](generics). كان لدينا في ذلك المثال
// أسلوب `AllElements` يعيد شريحة بجميع عناصر القائمة.
// يمكننا تنفيذ ذلك بصورة أفضل باستخدام مكررات Go، كما
// هو موضح أدناه.
type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

// تعيد `All` _مكررًا_، وهو في Go دالة ذات
// [توقيع خاص](https://pkg.go.dev/iter#Seq).
func (lst *List[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		// تستقبل دالة المكرر دالة أخرى كمُعامِل تسمى `yield`
		// حسب العرف، مع إمكان اختيار أي اسم. تستدعي `yield`
		// لكل عنصر نريد اجتيازه، وتراعي قيمة إرجاع `yield`
		// لاحتمال الإنهاء المبكر.
		for e := lst.head; e != nil; e = e.next {
			if !yield(e.val) {
				return
			}
		}
	}
}

// لا يتطلب الاجتياز هيكل بيانات أساسيًا، ولا يلزم حتى أن
// يكون محدودًا! إليك دالة تعيد مكررًا لأعداد فيبوناتشي؛
// إذ تستمر في العمل ما دامت `yield` تعيد `true`.
func genFib() iter.Seq[int] {
	return func(yield func(int) bool) {
		a, b := 0, 1

		for {
			if !yield(a) {
				return
			}
			a, b = b, a+b
		}
	}
}

func main() {
	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)

	// بما أن `List.All` تعيد مكررًا، يمكننا استخدامها في
	// حلقة `range` عادية.
	for e := range lst.All() {
		fmt.Println(e)
	}

	// تحتوي حزم مثل [slices](https://pkg.go.dev/slices) على
	// عدد من الدوال المفيدة للعمل مع المكررات. فمثلًا، تستقبل
	// `Collect` أي مكرر وتجمع كل قيمه في شريحة.
	all := slices.Collect(lst.All())
	fmt.Println("all:", all)

	// توفر حزم المكتبة القياسية الآن دوال مساعدة للمكررات
	// أيضًا. فمثلًا، تجتاز `strings.SplitSeq` أجزاء شريحة
	// بايتات دون إنشاء شريحة نتائج أولًا.
	for part := range strings.SplitSeq("go-by-example", "-") {
		fmt.Printf("part: %s\n", part)
	}

	for n := range genFib() {

		// بمجرد وصول الحلقة إلى `break` أو إرجاع مبكر، ستعيد
		// الدالة `yield` الممررة إلى المكرر القيمة `false`.
		if n >= 10 {
			break
		}
		fmt.Println(n)
	}
}
