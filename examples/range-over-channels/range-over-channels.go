// رأينا في مثال [سابق](range-over-built-in-types) كيف تتيح
// `for` و`range` اجتياز هياكل البيانات الأساسية. يمكننا أيضًا
// استخدام هذه الصياغة لاجتياز القيم المستقبلة من قناة.

package main

import "fmt"

func main() {

	// سنجتاز قيمتين في القناة `queue`.
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	// تجتاز `range` كل عنصر عند استقباله من `queue`. ولأننا أغلقنا
	// القناة أعلاه باستخدام `close`، ينتهي الاجتياز بعد استقبال
	// العنصرين.
	for elem := range queue {
		fmt.Println(elem)
	}
}
