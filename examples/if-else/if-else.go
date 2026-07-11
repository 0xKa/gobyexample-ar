// التفرع باستخدام `if` و`else` في Go مباشر وبسيط.

package main

import "fmt"

func main() {

	// إليك مثالًا أساسيًا.
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// يمكن استخدام عبارة `if` دون `else`.
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// غالبًا ما تكون المعاملات المنطقية مثل `&&` و`||`
	// مفيدة في الشروط.
	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("either 8 or 7 are even")
	}

	// يمكن أن تسبق الشروط عبارة؛ وتتاح المتغيرات المصرّح
	// عنها في هذه العبارة داخل الفرع الحالي وجميع الفروع
	// اللاحقة.
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}

// لاحظ أنك لا تحتاج إلى أقواس حول الشروط في Go، لكن
// الأقواس المعقوفة مطلوبة.
