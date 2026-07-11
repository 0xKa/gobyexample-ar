// توفر مكتبة Go القياسية دعمًا ممتازًا لعملاء HTTP وخوادم HTTP في
// الحزمة `net/http`. سنستخدمها في هذا المثال لإرسال طلبات HTTP
// بسيطة.
package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {

	// أرسل طلب HTTP من النوع `GET` إلى خادم. تمثل `http.Get` اختصارًا
	// ملائمًا يغني عن إنشاء كائن `http.Client` واستدعاء أسلوبه
	// `Get`؛ فهي تستخدم الكائن `http.DefaultClient` ذا الإعدادات
	// الافتراضية المفيدة.
	resp, err := http.Get("https://gobyexample.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// اطبع حالة استجابة HTTP.
	fmt.Println("Response status:", resp.Status)

	// اطبع أول 5 أسطر من جسم الاستجابة.
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
