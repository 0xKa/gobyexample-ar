// تسهل كتابة خادم HTTP أساسي باستخدام الحزمة `net/http`.
package main

import (
	"fmt"
	"net/http"
)

// تُعد *المعالجات* مفهومًا أساسيًا في خوادم `net/http`. المعالج
// كائن ينفذ الواجهة `http.Handler`. من الطرق الشائعة لكتابة معالج
// استخدام المهايئ `http.HandlerFunc` مع دوال ذات توقيع مناسب.
func hello(w http.ResponseWriter, req *http.Request) {

	// تأخذ الدوال التي تعمل كمعالجات `http.ResponseWriter`
	// و`http.Request` وسيطين. يُستخدم كاتب الاستجابة لملء استجابة
	// HTTP. استجابتنا البسيطة هنا ليست سوى `"hello\n"`.
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

	// ينفذ هذا المعالج مهمة أكثر تقدمًا قليلًا، إذ يقرأ جميع ترويسات
	// طلب HTTP ويرددها في جسم الاستجابة.
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {

	// نسجّل معالجاتنا في مسارات الخادم باستخدام الدالة المساعدة
	// `http.HandleFunc`. تضبط هذه الدالة *الموجّه الافتراضي* في
	// الحزمة `net/http`، وتأخذ دالة بوصفها وسيطًا.
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	// نستدعي أخيرًا `ListenAndServe` مع المنفذ ومعالج. تطلب منها
	// `nil` استخدام الموجّه الافتراضي الذي أعددناه للتو.
	http.ListenAndServe(":8090", nil)
}
