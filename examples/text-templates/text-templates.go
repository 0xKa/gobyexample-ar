// توفر Go دعمًا مدمجًا لإنشاء محتوى ديناميكي أو عرض خرج مخصص
// للمستخدم باستخدام الحزمة `text/template`. توفر الحزمة الشقيقة
// `html/template` الواجهة البرمجية نفسها، لكنها تتمتع بمزايا أمان
// إضافية وينبغي استخدامها لإنشاء HTML.

package main

import (
	"os"
	"text/template"
)

func main() {

	// يمكننا إنشاء قالب جديد وتحليل محتواه من سلسلة نصية.
	// تجمع القوالب بين نص ثابت و«إجراءات» محاطة بـ`{{...}}`، تُستخدم
	// لإدراج المحتوى ديناميكيًا.
	t1 := template.New("t1")
	t1, err := t1.Parse("Value is {{.}}\n")
	if err != nil {
		panic(err)
	}

	// بدلًا من ذلك، يمكننا استخدام الدالة `template.Must` لاستدعاء
	// `panic` إذا أعادت `Parse` خطأ. يفيد ذلك خصوصًا مع القوالب
	// المهيأة في النطاق العام.
	t1 = template.Must(t1.Parse("Value: {{.}}\n"))

	// «بتنفيذ» القالب، نُنشئ نصه باستخدام قيم محددة لإجراءاته.
	// يُستبدل الإجراء `{{.}}` بالقيمة الممررة كمُعامِل إلى
	// `Execute`.
	t1.Execute(os.Stdout, "some text")
	t1.Execute(os.Stdout, 5)
	t1.Execute(os.Stdout, []string{
		"Go",
		"Rust",
		"C++",
		"C#",
	})

	// دالة مساعدة سنستخدمها أدناه.
	Create := func(name, t string) *template.Template {
		return template.Must(template.New(name).Parse(t))
	}

	// إذا كانت البيانات هيكلًا، فيمكننا استخدام الإجراء
	// `{{.FieldName}}` للوصول إلى حقوله. يجب أن تكون الحقول مصدّرة
	// حتى يمكن الوصول إليها أثناء تنفيذ القالب.
	t2 := Create("t2", "Name: {{.Name}}\n")

	t2.Execute(os.Stdout, struct {
		Name string
	}{"Jane Doe"})

	// ينطبق الأمر نفسه على الخرائط، لكن لا قيود فيها على حالة أحرف
	// أسماء المفاتيح.
	t2.Execute(os.Stdout, map[string]string{
		"Name": "Mickey Mouse",
	})

	// توفر `if/else` تنفيذًا شرطيًا للقوالب. تُعد القيمة `false` إذا
	// كانت القيمة الافتراضية لنوع، مثل 0 أو سلسلة نصية فارغة أو
	// مؤشر `nil`، وما إلى ذلك. توضح هذه العينة ميزة أخرى للقوالب:
	// استخدام `-` في الإجراءات لإزالة المسافات البيضاء.
	t3 := Create("t3",
		"{{if . -}} yes {{else -}} no {{end}}\n")
	t3.Execute(os.Stdout, "not empty")
	t3.Execute(os.Stdout, "")

	// تتيح لنا كتل `range` اجتياز الشرائح أو المصفوفات أو الخرائط
	// أو القنوات. داخل كتلة `range`، تُضبط `{{.}}` على العنصر الحالي
	// من الاجتياز.
	t4 := Create("t4",
		"Range: {{range .}}{{.}} {{end}}\n")
	t4.Execute(os.Stdout,
		[]string{
			"Go",
			"Rust",
			"C++",
			"C#",
		})
}
