// توفر Go دعمًا مدمجًا
// [للتعبيرات النمطية](https://ar.wikipedia.org/wiki/تعبير_نمطي).
// إليك بعض الأمثلة على المهام الشائعة المتعلقة بها في Go.

package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {

	// يختبر هذا ما إذا كان النمط يطابق سلسلة نصية.
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	// استخدمنا أعلاه نمطًا نصيًا مباشرة، لكنك ستحتاج في مهام
	// التعبيرات النمطية الأخرى إلى استخدام `Compile` لترجمة النمط
	// إلى هيكل `Regexp` محسّن.
	r, _ := regexp.Compile("p([a-z]+)ch")

	// تتوفر أساليب كثيرة على هذه الهياكل. هذا اختبار تطابق شبيه
	// بما رأيناه سابقًا.
	fmt.Println(r.MatchString("peach"))

	// يجد هذا النص المطابق للتعبير النمطي.
	fmt.Println(r.FindString("peach punch"))

	// يجد هذا أيضًا أول تطابق، لكنه يعيد فهرسي بداية التطابق
	// ونهايته بدلًا من النص المطابق.
	fmt.Println("idx:", r.FindStringIndex("peach punch"))

	// تتضمن صيغ `Submatch` معلومات عن تطابق النمط الكامل
	// والتطابقات الفرعية داخله. سيعيد هذا مثلًا معلومات عن كل من
	// `p([a-z]+)ch` و`([a-z]+)`.
	fmt.Println(r.FindStringSubmatch("peach punch"))

	// وبالمثل، سيعيد هذا معلومات عن فهارس التطابقات والتطابقات
	// الفرعية.
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))

	// تنطبق صيغ `All` من هذه الدوال على جميع التطابقات في
	// المدخلات، وليس أول تطابق فقط. يجد هذا مثلًا جميع التطابقات
	// مع تعبير نمطي.
	fmt.Println(r.FindAllString("peach punch pinch", -1))

	// تتوفر صيغ `All` هذه أيضًا للدوال الأخرى التي رأيناها
	// أعلاه.
	fmt.Println("all:", r.FindAllStringSubmatchIndex(
		"peach punch pinch", -1))

	// يؤدي تمرير عدد صحيح غير سالب وسيطًا ثانيًا لهذه الدوال إلى
	// تقييد عدد التطابقات.
	fmt.Println(r.FindAllString("peach punch pinch", 2))

	// استخدمت أمثلتنا أعلاه وسائط من السلاسل النصية وأسماء مثل
	// `MatchString`. يمكننا أيضًا تمرير وسائط من النوع `[]byte`
	// وحذف `String` من اسم الدالة.
	fmt.Println(r.Match([]byte("peach")))

	// عند إنشاء متغيرات عامة تحتوي على تعبيرات نمطية، يمكنك
	// استخدام الصيغة `MustCompile` من `Compile`. تستدعي
	// `MustCompile` الدالة `panic` بدلًا من إعادة خطأ، ما يجعل
	// استخدامها أكثر أمانًا للمتغيرات العامة.
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println("regexp:", r)

	// يمكن أيضًا استخدام الحزمة `regexp` لاستبدال أجزاء من
	// السلاسل النصية بقيم أخرى.
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	// تتيح لك الصيغة `Func` تحويل النص المطابق باستخدام دالة
	// محددة.
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}
