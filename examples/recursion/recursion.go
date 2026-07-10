// تدعم Go
// <a href="https://en.wikipedia.org/wiki/Recursion_(computer_science)"><em>الدوال ذاتية الاستدعاء</em></a>.
// إليك مثالًا تقليديًا.

package main

import "fmt"

// تستدعي الدالة `fact` نفسها حتى تصل إلى الحالة الأساسية
// `fact(0)`.
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(7))

	// يمكن للدوال المجهولة أيضًا أن تستدعي نفسها، لكن ذلك
	// يتطلب التصريح صراحةً باستخدام `var` عن متغير يخزن
	// الدالة قبل تعريفها.
	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}

		// بما أن `fib` صُرّح عنها سابقًا في `main`، تعرف Go
		// أي دالة تستدعيها باستخدام `fib` هنا.
		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(7))
}
