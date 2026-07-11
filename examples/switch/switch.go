// تعبّر عبارات `switch` عن الشروط التي تتوزع على عدة
// فروع.

package main

import (
	"fmt"
	"time"
)

func main() {

	// إليك عبارة `switch` أساسية.
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// يمكنك استخدام الفواصل للفصل بين عدة تعبيرات في عبارة
	// `case` نفسها. ونستخدم في هذا المثال أيضًا حالة
	// `default` الاختيارية.
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	// تمثل `switch` دون تعبير طريقة بديلة لصياغة منطق
	// `if/else`. ونوضح هنا أيضًا أن تعبيرات `case` يمكن
	// أن تكون غير ثابتة.
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	// تقارن `switch` الخاصة بالأنواع بين الأنواع بدلًا من
	// القيم. يمكنك استخدامها لمعرفة نوع قيمة مخزنة في واجهة.
	// في هذا المثال، سيحمل المتغير `t` النوع الموافق لفرعه.
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
