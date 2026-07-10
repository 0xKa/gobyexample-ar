// عند استخدام القنوات كمُعامِلات للدوال، يمكنك تحديد ما إذا
// كانت القناة مخصصة لإرسال القيم فقط أو لاستقبالها فقط. يزيد
// هذا التحديد من أمان الأنواع في البرنامج.

package main

import "fmt"

// لا تقبل الدالة `ping` سوى قناة لإرسال القيم. ستؤدي محاولة
// الاستقبال من هذه القناة إلى خطأ في وقت الترجمة.
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// تقبل الدالة `pong` قناة للاستقبال (`pings`) وأخرى للإرسال
// (`pongs`).
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
