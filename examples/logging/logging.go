// توفر مكتبة Go القياسية أدوات مباشرة لإخراج السجلات من برامج
// Go، باستخدام الحزمة [log](https://pkg.go.dev/log) للخرج الحر،
// والحزمة [log/slog](https://pkg.go.dev/log/slog) للخرج المنظم.
package main

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"log/slog"
)

func main() {

	// يستخدم مجرد استدعاء دوال مثل `Println` من الحزمة `log`
	// مسجّل الأحداث _القياسي_، المهيأ مسبقًا لإخراج سجلات مناسبة
	// إلى `os.Stderr`. تؤدي أساليب إضافية مثل `Fatal*` أو `Panic*`
	// إلى خروج البرنامج بعد التسجيل.
	log.Println("standard logger")

	// يمكن إعداد مسجّلات الأحداث باستخدام _خيارات_ تحدد صيغة
	// خرجها. تكون خيارات `log.Ldate` و`log.Ltime` مفعلة افتراضيًا
	// في المسجّل القياسي، وهي مجمعة في `log.LstdFlags`. يمكننا مثلًا
	// تغيير خياراته لإخراج الوقت بدقة الميكروثانية.
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("with micro")

	// يدعم أيضًا إخراج اسم الملف ورقم السطر الذي استُدعيت منه دالة
	// `log`.
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("with file/line")

	// قد يفيد إنشاء مسجّل أحداث مخصص وتمريره. عند إنشاء مسجّل جديد،
	// يمكننا تعيين _بادئة_ لتمييز خرجه من خرج المسجّلات الأخرى.
	mylog := log.New(os.Stdout, "my:", log.LstdFlags)
	mylog.Println("from mylog")

	// يمكننا تعيين البادئة لمسجّلات الأحداث الموجودة، بما فيها
	// المسجّل القياسي، باستخدام الأسلوب `SetPrefix`.
	mylog.SetPrefix("ohmy:")
	mylog.Println("from mylog")

	// يمكن أن تكون لمسجّلات الأحداث أهداف خرج مخصصة؛ إذ يصلح أي
	// `io.Writer`.
	var buf bytes.Buffer
	buflog := log.New(&buf, "buf:", log.LstdFlags)

	// يكتب هذا الاستدعاء خرج السجل في `buf`.
	buflog.Println("hello")

	// سيعرضه هذا فعليًا على الخرج القياسي.
	fmt.Print("from buflog:", buf.String())

	// توفر الحزمة `slog` خرج سجلات _منظمًا_. ويكون تسجيل الأحداث
	// بصيغة JSON مباشرًا مثلًا.
	jsonHandler := slog.NewJSONHandler(os.Stderr, nil)
	myslog := slog.New(jsonHandler)
	myslog.Info("hi there")

	// يمكن أن يحتوي خرج `slog`، إضافة إلى الرسالة، على أي عدد من
	// أزواج `key=value`.
	myslog.Info("hello again", "key", "val", "age", 25)
}
