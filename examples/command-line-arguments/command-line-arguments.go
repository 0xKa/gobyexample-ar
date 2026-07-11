// تُعد [وسائط سطر الأوامر](https://ar.wikipedia.org/wiki/واجهة_سطر_الأوامر)
// طريقة شائعة لتحديد مُعامِلات تشغيل البرامج. يستخدم الأمر
// `go run hello.go` مثلًا `run` و`hello.go` وسيطين للبرنامج `go`.

package main

import (
	"fmt"
	"os"
)

func main() {

	// تتيح `os.Args` الوصول إلى وسائط سطر الأوامر الخام. لاحظ أن
	// أول قيمة في هذه الشريحة هي مسار البرنامج، وأن `os.Args[1:]`
	// تحتوي على وسائط البرنامج.
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	// يمكنك الحصول على كل وسيط على حدة باستخدام الفهرسة المعتادة.
	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}
