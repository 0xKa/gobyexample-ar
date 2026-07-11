// تحتاج برامج Go أحيانًا إلى تشغيل عمليات أخرى.

package main

import (
	"errors"
	"fmt"
	"io"
	"os/exec"
)

func main() {

	// سنبدأ بأمر بسيط لا يأخذ وسائط أو مدخلات، ويطبع شيئًا إلى
	// `stdout` فحسب. تنشئ الدالة المساعدة `exec.Command` كائنًا
	// يمثل هذه العملية الخارجية.
	dateCmd := exec.Command("date")

	// يشغّل الأسلوب `Output` الأمر وينتظر انتهاءه ويجمع خرجه
	// القياسي. إذا لم تحدث أخطاء، فستحتوي `dateOut` على بايتات
	// معلومات التاريخ.
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(dateOut))

	// يعيد `Output` وأساليب `Command` الأخرى قيمة `*exec.Error` إذا
	// حدثت مشكلة أثناء تنفيذ الأمر، مثل مسار خاطئ، ويعيد
	// `*exec.ExitError` إذا عمل الأمر لكنه خرج برمز إرجاع غير صفري.
	_, err = exec.Command("date", "-x").Output()
	if err != nil {
		if e, ok := errors.AsType[*exec.Error](err); ok {
			fmt.Println("failed executing:", e)
		} else if e, ok := errors.AsType[*exec.ExitError](err); ok {
			exitCode := e.ExitCode()
			fmt.Println("command exit rc =", exitCode)
		} else {
			panic(err)
		}
	}

	// سنتعرف تاليًا على حالة أعقد قليلًا، نمرر فيها البيانات عبر
	// أنبوب إلى `stdin` للعملية الخارجية، ونجمع النتائج من
	// `stdout` الخاص بها.
	grepCmd := exec.Command("grep", "hello")

	// نحصل هنا صراحة على أنبوبي الإدخال والإخراج، ونبدأ العملية،
	// ونكتب بعض المدخلات إليها، ونقرأ الخرج الناتج، ثم ننتظر أخيرًا
	// خروج العملية.
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()
	grepBytes, _ := io.ReadAll(grepOut)
	grepCmd.Wait()

	// أغفلنا فحوص الأخطاء في المثال أعلاه، لكن يمكنك استخدام النمط
	// المعتاد `if err != nil` معها كلها. كما أننا لا نجمع إلا نتائج
	// `StdoutPipe`، لكن يمكنك جمع `StderrPipe` بالطريقة نفسها تمامًا.
	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	// لاحظ أننا عند تشغيل الأوامر نحتاج إلى توفير الأمر ومصفوفة
	// الوسائط كلًا على حدة، بدلًا من تمرير سلسلة نصية واحدة لسطر
	// الأوامر. إذا أردت تشغيل أمر كامل من سلسلة نصية، فيمكنك استخدام
	// الخيار `-c` في `bash`:
	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))
}
