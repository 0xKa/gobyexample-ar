// قراءة الملفات وكتابتها مهمتان أساسيتان تحتاج إليهما برامج Go
// كثيرة. سنتعرف أولًا على بعض أمثلة قراءة الملفات.

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// تتطلب قراءة الملفات التحقق من الأخطاء في معظم الاستدعاءات.
// ستبسّط هذه الدالة المساعدة فحوص الأخطاء أدناه.
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// لعل أبسط مهام قراءة الملفات هي تحميل محتويات ملف كاملة إلى
	// الذاكرة.
	path := filepath.Join(os.TempDir(), "dat")
	dat, err := os.ReadFile(path)
	check(err)
	fmt.Print(string(dat))

	// ستحتاج غالبًا إلى تحكم أكبر في كيفية قراءة الملف والأجزاء
	// المقروءة منه. لهذه المهام، ابدأ بفتح ملف باستخدام `Open`
	// للحصول على قيمة `os.File`.
	f, err := os.Open(path)
	check(err)

	// اقرأ بعض البايتات من بداية الملف. اسمح بقراءة ما يصل إلى 5
	// بايتات، وسجّل أيضًا عدد ما قُرئ فعليًا.
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

	// يمكنك أيضًا الانتقال باستخدام `Seek` إلى موضع معلوم في الملف،
	// ثم القراءة منه باستخدام `Read`.
	o2, err := f.Seek(6, io.SeekStart)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))

	// توجد طرق أخرى للانتقال نسبة إلى موضع المؤشر الحالي،
	_, err = f.Seek(2, io.SeekCurrent)
	check(err)

	// ونسبة إلى نهاية الملف.
	_, err = f.Seek(-4, io.SeekEnd)
	check(err)

	// توفر الحزمة `io` بعض الدوال التي قد تفيد في قراءة الملفات.
	// يمكن مثلًا تنفيذ عمليات قراءة شبيهة بما سبق بمتانة أكبر
	// باستخدام `ReadAtLeast`.
	o3, err := f.Seek(6, io.SeekStart)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	// لا توجد عملية مدمجة لإعادة المؤشر إلى البداية، لكن
	// `Seek(0, io.SeekStart)` تحقق ذلك.
	_, err = f.Seek(0, io.SeekStart)
	check(err)

	// تنفذ الحزمة `bufio` قارئًا مخزنًا مؤقتًا قد يفيد لكفاءته مع
	// عمليات القراءة الصغيرة الكثيرة، ولما يوفره من أساليب قراءة
	// إضافية.
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	// أغلق الملف عند الانتهاء (عادة ما يُجدول ذلك باستخدام `defer`
	// فور فتح الملف عبر `Open`).
	f.Close()
}
