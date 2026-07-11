// من المتعارف عليه في Go التعبير عن الأخطاء بقيمة إرجاع
// صريحة ومنفصلة. يختلف ذلك عن الاستثناءات المستخدمة في
// لغات مثل Java وPython وRuby، وعن القيمة المفردة التي
// تُستخدم أحيانًا في C لتمثيل النتيجة أو الخطأ. تسهّل
// طريقة Go معرفة الدوال التي تعيد أخطاء ومعالجتها باستخدام
// تراكيب اللغة نفسها المستخدمة للمهام الأخرى غير المرتبطة
// بالأخطاء.
//
// راجع توثيق [حزمة errors](https://pkg.go.dev/errors)
// و[هذه المقالة](https://go.dev/blog/go1.13-errors) لمزيد
// من التفاصيل.

package main

import (
	"errors"
	"fmt"
)

// تكون الأخطاء حسب العرف قيمة الإرجاع الأخيرة، ويكون نوعها
// `error`، وهو واجهة مدمجة.
func f(arg int) (int, error) {
	if arg == 42 {
		// تنشئ `errors.New` قيمة `error` أساسية برسالة
		// الخطأ المعطاة.
		return -1, errors.New("can't work with 42")
	}

	// تدل قيمة `nil` في موضع الخطأ على عدم وقوع خطأ.
	return arg + 3, nil
}

// الخطأ الدالّ متغير مصرح عنه مسبقًا يُستخدم للدلالة على
// حالة خطأ محددة.
var ErrOutOfTea = errors.New("no more tea available")
var ErrPower = errors.New("can't boil water")

func makeTea(arg int) error {
	if arg == 2 {
		return ErrOutOfTea
	} else if arg == 4 {

		// يمكننا تغليف الأخطاء بأخطاء أعلى مستوى لإضافة سياق.
		// أبسط طريقة لذلك هي استخدام الفعل `%w` في
		// `fmt.Errorf`. تنشئ الأخطاء المغلفة سلسلة منطقية
		// (يغلف `A` الخطأ `B`، الذي يغلف `C`، وهكذا) يمكن الاستعلام
		// عنها بدوال مثل `errors.Is` و`errors.AsType`.
		return fmt.Errorf("making tea: %w", ErrPower)
	}
	return nil
}

func main() {
	for _, i := range []int{7, 42} {

		// من المتعارف عليه التحقق من الخطأ داخل سطر `if`.
		if r, e := f(i); e != nil {
			fmt.Println("f failed:", e)
		} else {
			fmt.Println("f worked:", r)
		}
	}

	for i := range 5 {
		if err := makeTea(i); err != nil {

			// تتحقق `errors.Is` مما إذا كان خطأ معطى، أو أي خطأ
			// في سلسلته، يطابق قيمة خطأ محددة. يفيد ذلك خصوصًا
			// مع الأخطاء المغلفة أو المتداخلة، إذ يتيح تحديد أنواع
			// أخطاء معينة أو أخطاء دالّة داخل سلسلة أخطاء.
			if errors.Is(err, ErrOutOfTea) {
				fmt.Println("We should buy new tea!")
			} else if errors.Is(err, ErrPower) {
				fmt.Println("Now it is dark.")
			} else {
				fmt.Printf("unknown error: %s\n", err)
			}
			continue
		}

		fmt.Println("Tea is ready!")
	}
}
