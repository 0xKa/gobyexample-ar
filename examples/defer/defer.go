// تُستخدم الكلمة _`defer`_ لضمان تنفيذ استدعاء دالة في وقت لاحق
// من تشغيل البرنامج، ويكون ذلك عادة لأغراض التنظيف. تُستخدم
// `defer` غالبًا حيث تُستخدم `ensure` و`finally` مثلًا في لغات
// أخرى.

package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// لنفترض أننا نريد إنشاء ملف والكتابة إليه، ثم إغلاقه عند
// الانتهاء. هكذا يمكننا فعل ذلك باستخدام `defer`.
func main() {

	// فور الحصول على كائن ملف باستخدام `createFile`، نؤجل إغلاق
	// الملف عبر `closeFile`. سيُنفذ هذا الاستدعاء عند نهاية الدالة
	// المحيطة (`main`)، بعد انتهاء `writeFile`.
	path := filepath.Join(os.TempDir(), "defer.txt")
	f := createFile(path)
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()
	// من المهم التحقق من الأخطاء عند إغلاق ملف، حتى داخل دالة مؤجلة.
	if err != nil {
		panic(err)
	}
}
