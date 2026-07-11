// تعرفنا في المثال السابق على [تشغيل عمليات
// خارجية](spawning-processes). نفعل ذلك عندما نحتاج إلى عملية
// خارجية تستطيع عملية Go قيد التشغيل الوصول إليها. نريد أحيانًا
// استبدال عملية Go الحالية بالكامل بعملية أخرى، قد لا تكون مكتوبة
// بلغة Go. لفعل ذلك، سنستخدم تنفيذ Go لدالة
// <a href="https://en.wikipedia.org/wiki/Exec_(operating_system)"><code>exec</code></a>
// التقليدية.

package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main() {

	// سنستبدل عمليتنا في هذا المثال بـ`ls`. تتطلب Go مسارًا مطلقًا
	// للملف التنفيذي الذي نريد تشغيله، لذا سنستخدم `exec.LookPath`
	// للعثور عليه، ويُرجح أن يكون `/bin/ls`.
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	// تتطلب `Exec` الوسائط في صورة شريحة، لا سلسلة نصية كبيرة واحدة.
	// سنمرر إلى `ls` بعض الوسائط الشائعة. لاحظ أن الوسيط الأول ينبغي
	// أن يكون اسم البرنامج.
	args := []string{"ls", "-a", "-l", "-h"}

	// تحتاج `Exec` أيضًا إلى مجموعة من [متغيرات
	// البيئة](environment-variables) لاستخدامها. نوفر هنا بيئتنا
	// الحالية فحسب.
	env := os.Environ()

	// هذا هو استدعاء `syscall.Exec` الفعلي. إذا نجح الاستدعاء،
	// فسينتهي تنفيذ عمليتنا هنا وتحل محله العملية
	// `/bin/ls -a -l -h`. وإذا حدث خطأ، فسنحصل على قيمة إرجاع.
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
