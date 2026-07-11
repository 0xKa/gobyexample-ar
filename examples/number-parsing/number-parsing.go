// تحليل الأعداد من السلاسل النصية مهمة أساسية وشائعة في كثير من
// البرامج. هكذا تفعل ذلك في Go.

package main

// توفر الحزمة `strconv` في المكتبة القياسية تحليل الأعداد.
import (
	"fmt"
	"strconv"
)

func main() {

	// مع `ParseFloat`، تحدد `64` عدد بتات الدقة المطلوب تحليلها.
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	// مع `ParseInt`، تعني `0` استنتاج الأساس من السلسلة النصية،
	// وتشترط `64` أن تتسع النتيجة في 64 بتًا.
	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	// تتعرف `ParseInt` على الأعداد المنسقة بالنظام السداسي عشري.
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	// تتوفر أيضًا `ParseUint`.
	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	// `Atoi` دالة مساعدة للتحليل الأساسي لقيم `int` ذات الأساس 10.
	k, _ := strconv.Atoi("135")
	fmt.Println(k)

	// تعيد دوال التحليل خطأ عند تلقي مدخلات غير صالحة.
	_, e := strconv.Atoi("wat")
	fmt.Println(e)
}
