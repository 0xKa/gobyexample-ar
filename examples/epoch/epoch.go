// من المتطلبات الشائعة في البرامج الحصول على عدد الثواني أو
// المللي ثانية أو النانو ثانية المنقضية منذ
// [حقبة Unix](https://ar.wikipedia.org/wiki/توقيت_يونكس).
// هكذا تفعل ذلك في Go.

package main

import (
	"fmt"
	"time"
)

func main() {

	// استخدم `time.Now` مع `Unix` أو `UnixMilli` أو `UnixNano`
	// للحصول على الوقت المنقضي منذ حقبة Unix بالثواني أو المللي
	// ثانية أو النانو ثانية على الترتيب.
	now := time.Now()
	fmt.Println(now)

	fmt.Println(now.Unix())
	fmt.Println(now.UnixMilli())
	fmt.Println(now.UnixNano())

	// يمكنك أيضًا تحويل عدد صحيح من الثواني أو النانو ثانية منذ
	// الحقبة إلى قيمة `time` المقابلة.
	fmt.Println(time.Unix(now.Unix(), 0))
	fmt.Println(time.Unix(0, now.UnixNano()))
}
