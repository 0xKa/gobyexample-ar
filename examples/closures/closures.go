// تدعم Go [_الدوال المجهولة_](https://en.wikipedia.org/wiki/Anonymous_function)،
// التي يمكنها تكوين <a href="https://en.wikipedia.org/wiki/Closure_(computer_science)"><em>إغلاقات</em></a>.
// تفيد الدوال المجهولة عندما تريد تعريف دالة في موضعها
// دون الحاجة إلى تسميتها.

package main

import "fmt"

// تعيد الدالة `intSeq` دالة أخرى نعرّفها دون اسم داخل
// جسم `intSeq`. تلتقط الدالة المعادة المتغير `i` لتكوّن
// إغلاقًا.
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	// نستدعي `intSeq` ونسند النتيجة، وهي دالة، إلى
	// `nextInt`. تلتقط هذه الدالة قيمة `i` الخاصة بها،
	// التي ستُحدَّث في كل مرة نستدعي فيها `nextInt`.
	nextInt := intSeq()

	// شاهد أثر الإغلاق باستدعاء `nextInt` عدة مرات.
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// للتأكد من أن الحالة خاصة بتلك الدالة، أنشئ دالة
	// جديدة واختبرها.
	newInts := intSeq()
	fmt.Println(newInts())
}
