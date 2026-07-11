// توفر Go دعمًا مدمجًا [لترميز Base64 وفك
// ترميزه](https://ar.wikipedia.org/wiki/الأساس_64).

package main

// تستورد هذه الصياغة الحزمة `encoding/base64` بالاسم `b64` بدلًا
// من الاسم الافتراضي `base64`، ما يوفر علينا بعض المساحة أدناه.
import (
	b64 "encoding/base64"
	"fmt"
)

func main() {

	// هذه هي قيمة `string` التي سنرمّزها ونفك ترميزها.
	data := "abc123!?$*&()'-=@~"

	// تدعم Go صيغة Base64 القياسية والمتوافقة مع URL. هكذا نُجري
	// الترميز باستخدام المرمّز القياسي. يتطلب المرمّز قيمة
	// `[]byte`، لذلك نحوّل قيمة `string` لدينا إلى ذلك النوع.
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	// قد يعيد فك الترميز خطأ، ويمكنك التحقق منه إذا لم تكن تعرف
	// مسبقًا أن المدخلات سليمة الصياغة.
	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))
	fmt.Println()

	// يُجري هذا الترميز وفك الترميز باستخدام صيغة Base64 متوافقة
	// مع URL.
	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))
}
