// `//go:embed` هو [توجيه
// للمترجم](https://pkg.go.dev/cmd/compile#hdr-Compiler_Directives) يتيح
// للبرامج تضمين أي ملفات ومجلدات في ملف Go التنفيذي عند وقت البناء.
// اقرأ المزيد عن توجيه `embed`
// [هنا](https://pkg.go.dev/embed).
package main

// استورد الحزمة `embed`. إذا لم تستخدم أي معرّفات مصدّرة من هذه
// الحزمة، فيمكنك إجراء استيراد فارغ باستخدام `_ "embed"`.
import (
	"embed"
)

// تقبل توجيهات `embed` مسارات نسبية إلى المجلد الذي يحتوي على ملف
// Go المصدري. يضمّن هذا التوجيه محتويات الملف في متغير `string`
// الذي يليه مباشرة.
//
//go:embed folder/single_file.txt
var fileString string

// أو ضمّن محتويات الملف في قيمة `[]byte`.
//
//go:embed folder/single_file.txt
var fileByte []byte

// يمكننا أيضًا تضمين عدة ملفات، أو حتى مجلدات، باستخدام محارف
// البدل. يستخدم ذلك متغيرًا من [النوع embed.FS](https://pkg.go.dev/embed#FS)،
// الذي ينفذ نظام ملفات افتراضيًا بسيطًا.
//
//go:embed folder/single_file.txt
//go:embed folder/*.hash
var folder embed.FS

func main() {

	// اطبع محتويات `single_file.txt`.
	print(fileString)
	print(string(fileByte))

	// استرجع بعض الملفات من المجلد المضمّن.
	content1, _ := folder.ReadFile("folder/file1.hash")
	print(string(content1))

	content2, _ := folder.ReadFile("folder/file2.hash")
	print(string(content2))
}
