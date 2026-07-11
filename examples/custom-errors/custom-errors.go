// يمكن تعريف أنواع أخطاء مخصصة بتطبيق الأسلوب `Error()`
// عليها. إليك نسخة من المثال السابق تستخدم نوعًا مخصصًا
// لتمثيل خطأ في وسيط بوضوح.

package main

import (
	"errors"
	"fmt"
)

// يحمل نوع الخطأ المخصص عادةً اللاحقة `Error`.
type argError struct {
	arg     int
	message string
}

// تؤدي إضافة الأسلوب `Error` إلى جعل `argError` يطبق
// واجهة `error`.
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.message)
}

func f(arg int) (int, error) {
	if arg == 42 {

		// أعد الخطأ المخصص.
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {

	// `errors.AsType` نسخة أكثر تقدمًا من `errors.Is`. تتحقق
	// مما إذا كان خطأ معطى، أو أي خطأ في سلسلته، يطابق نوع
	// خطأ محدد، وتحوله إلى قيمة من ذلك النوع مع إعادة
	// `true` أيضًا. إذا لم يوجد تطابق، تكون قيمة الإرجاع
	// الثانية `false`.
	_, err := f(42)
	if ae, ok := errors.AsType[*argError](err); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.message)
	} else {
		fmt.Println("err doesn't match argError")
	}
}
