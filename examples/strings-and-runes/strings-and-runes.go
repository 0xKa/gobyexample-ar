// السلسلة النصية في Go شريحة بايتات للقراءة فقط. تتعامل
// اللغة والمكتبة القياسية مع السلاسل النصية معاملة خاصة،
// بوصفها حاويات لنص مرمّز بصيغة [UTF-8](https://en.wikipedia.org/wiki/UTF-8).
// تتكون السلاسل النصية في لغات أخرى من «محارف». أما في
// Go، فيسمى مفهوم المحرف `rune`، وهو عدد صحيح يمثل نقطة
// ترميز Unicode. تقدم [هذه المقالة في مدونة Go](https://go.dev/blog/strings)
// مقدمة جيدة للموضوع.

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	// `s` سلسلة نصية من النوع `string` أُسندت إليها قيمة
	// حرفية تمثل كلمة «مرحبًا» باللغة التايلاندية. القيم
	// الحرفية للسلاسل النصية في Go نص مرمّز بصيغة UTF-8.
	const s = "สวัสดี"

	// بما أن السلاسل النصية تكافئ `[]byte`، فسيعطي هذا
	// طول البايتات الخام المخزنة داخلها.
	fmt.Println("Len:", len(s))

	// يعطي الفهرسة داخل سلسلة نصية قيم البايتات الخام عند
	// كل فهرس. تولد هذه الحلقة القيم الست عشرية لجميع
	// البايتات التي تكوّن نقاط الترميز في `s`.
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	// لعد محارف _`rune`_ الموجودة في سلسلة نصية، يمكننا
	// استخدام الحزمة `utf8`. لاحظ أن زمن تنفيذ
	// `RuneCountInString` يعتمد على حجم السلسلة النصية، لأنها
	// تفك ترميز كل محرف UTF-8 بالتتابع. تُمثّل بعض المحارف
	// التايلاندية بنقاط ترميز UTF-8 تمتد عبر عدة بايتات، لذلك
	// قد تكون نتيجة هذا العد مفاجئة.
	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	// تتعامل حلقة `range` مع السلاسل النصية معاملة خاصة،
	// فتفك ترميز كل محرف `rune` مع إزاحته في السلسلة.
	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	// يمكننا تنفيذ الاجتياز نفسه باستخدام الدالة
	// `utf8.DecodeRuneInString` صراحةً.
	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		// يوضح هذا تمرير قيمة من النوع `rune` إلى دالة.
		examineRune(runeValue)
	}
}

func examineRune(r rune) {

	// القيم المحاطة بعلامتي اقتباس مفردتين هي _قيم حرفية من
	// النوع `rune`_. ويمكننا مقارنة قيمة `rune` بقيمة حرفية
	// من النوع نفسه مباشرةً.
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}
