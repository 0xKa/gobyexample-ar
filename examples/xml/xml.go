// توفر Go دعمًا مدمجًا لـXML والصيغ الشبيهة به باستخدام الحزمة
// `encoding/xml`.

package main

import (
	"encoding/xml"
	"fmt"
)

// سيُمثّل `Plant` في XML. كما في أمثلة JSON، تحتوي وسوم الحقول
// على توجيهات للمرمّز وأداة فك الترميز. نستخدم هنا بعض مزايا حزمة
// XML الخاصة: يحدد اسم الحقل `XMLName` اسم عنصر XML الذي يمثل
// هذا الهيكل، وتعني `id,attr` أن الحقل `Id` هو _سمة_ XML وليس
// عنصرًا متداخلًا.
type Plant struct {
	XMLName xml.Name `xml:"plant"`
	Id      int      `xml:"id,attr"`
	Name    string   `xml:"name"`
	Origin  []string `xml:"origin"`
}

func (p Plant) String() string {
	return fmt.Sprintf("Plant id=%v, name=%v, origin=%v",
		p.Id, p.Name, p.Origin)
}

func main() {
	coffee := &Plant{Id: 27, Name: "Coffee"}
	coffee.Origin = []string{"Ethiopia", "Brazil"}

	// أخرج XML الذي يمثل نبتتنا، باستخدام `MarshalIndent` لإنتاج
	// خرج أسهل للقراءة.
	out, _ := xml.MarshalIndent(coffee, " ", "  ")
	fmt.Println(string(out))

	// لإضافة ترويسة XML عامة إلى الخرج، ألحقها صراحة.
	fmt.Println(xml.Header + string(out))

	// استخدم `Unmarshal` لتحليل تدفق بايتات يحتوي على XML إلى هيكل
	// بيانات. إذا كان XML غير صالح أو تعذر مطابقته مع `Plant`،
	// فسيُعاد خطأ وصفي.
	var p Plant
	if err := xml.Unmarshal(out, &p); err != nil {
		panic(err)
	}
	fmt.Println(p)

	tomato := &Plant{Id: 81, Name: "Tomato"}
	tomato.Origin = []string{"Mexico", "California"}

	// يطلب وسم الحقل `parent>child>plant` من المرمّز وضع جميع عناصر
	// `plant` متداخلة تحت `<parent><child>...`.
	type Nesting struct {
		XMLName xml.Name `xml:"nesting"`
		Plants  []*Plant `xml:"parent>child>plant"`
	}

	nesting := &Nesting{}
	nesting.Plants = []*Plant{coffee, tomato}

	out, _ = xml.MarshalIndent(nesting, " ", "  ")
	fmt.Println(string(out))
}
