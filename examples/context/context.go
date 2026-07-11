// تعرفنا في المثال السابق على إعداد [خادم HTTP](http-server)
// بسيط. تفيد خوادم HTTP في توضيح استخدام `context.Context`
// للتحكم في الإلغاء. ينقل `Context` المواعيد النهائية وإشارات
// الإلغاء والقيم الأخرى ذات نطاق الطلب عبر حدود الواجهات
// البرمجية وروتينات Go.
package main

import (
	"fmt"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {

	// تنشئ آلية `net/http` قيمة `context.Context` لكل طلب، ويمكن
	// الوصول إليها باستخدام الأسلوب `Context()`.
	ctx := req.Context()
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler ended")

	// انتظر بضع ثوانٍ قبل إرسال رد إلى العميل. قد يحاكي ذلك بعض
	// العمل الذي ينفذه الخادم. أثناء العمل، راقب قناة السياق
	// `Done()` بحثًا عن إشارة تطلب إلغاء العمل والعودة في أسرع وقت.
	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "hello\n")
	case <-ctx.Done():
		// يعيد الأسلوب `Err()` للسياق خطأ يوضح سبب إغلاق قناة
		// `Done()`.
		err := ctx.Err()
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}

func main() {

	// كما سبق، نسجّل معالجنا في المسار `"/hello"` ونبدأ الخدمة.
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8090", nil)
}
