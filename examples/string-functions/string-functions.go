// توفر الحزمة `strings` في المكتبة القياسية كثيرًا من الدوال
// المفيدة المتعلقة بالسلاسل النصية. إليك بعض الأمثلة للتعرف على
// الحزمة.

package main

import (
	"fmt"
	s "strings"
)

// نمنح `fmt.Println` اسمًا مستعارًا أقصر لأننا سنستخدمها كثيرًا
// أدناه.
var p = fmt.Println

func main() {

	// هذه عينة من الدوال المتاحة في `strings`. ولأنها دوال من
	// الحزمة وليست أساليب معرّفة على قيمة السلسلة النصية نفسها،
	// فعلينا تمرير السلسلة المعنية بصفتها الوسيط الأول للدالة. ستجد
	// مزيدًا من الدوال في توثيق الحزمة
	// [`strings`](https://pkg.go.dev/strings).
	p("Contains:  ", s.Contains("test", "es"))
	p("Count:     ", s.Count("test", "t"))
	p("HasPrefix: ", s.HasPrefix("test", "te"))
	p("HasSuffix: ", s.HasSuffix("test", "st"))
	p("Index:     ", s.Index("test", "e"))
	p("Join:      ", s.Join([]string{"a", "b"}, "-"))
	p("Repeat:    ", s.Repeat("a", 5))
	p("Replace:   ", s.Replace("foo", "o", "0", -1))
	p("Replace:   ", s.Replace("foo", "o", "0", 1))
	p("Split:     ", s.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", s.ToLower("TEST"))
	p("ToUpper:   ", s.ToUpper("test"))
}
