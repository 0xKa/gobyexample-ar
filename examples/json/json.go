// توفر Go دعمًا مدمجًا لترميز JSON وفك ترميزه، بما في ذلك
// التحويل من أنواع البيانات المدمجة والمخصصة وإليها.

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// سنستخدم هذين الهيكلين أدناه لتوضيح ترميز الأنواع المخصصة وفك
// ترميزها.
type response1 struct {
	Page   int
	Fruits []string
}

// لا تخضع للترميز وفك الترميز في JSON إلا الحقول المصدّرة. يجب
// أن تبدأ أسماء الحقول بأحرف إنجليزية كبيرة لتكون مصدّرة.
type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {

	// سنتعرف أولًا على ترميز أنواع البيانات الأساسية إلى سلاسل
	// نصية بصيغة JSON. إليك بعض الأمثلة على القيم البسيطة.
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	// وإليك أمثلة على الشرائح والخرائط، التي تُرمّز إلى مصفوفات
	// وكائنات JSON كما هو متوقع.
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	// يمكن لحزمة JSON ترميز أنواع بياناتك المخصصة تلقائيًا. لن
	// يتضمن الخرج المرمّز سوى الحقول المصدّرة، وستستخدم الحزمة
	// أسماءها افتراضيًا كمفاتيح JSON.
	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	// يمكنك استخدام الوسوم في تصريحات حقول الهيكل لتخصيص أسماء
	// مفاتيح JSON المرمّزة. راجع تعريف `response2` أعلاه للاطلاع
	// على مثال لهذه الوسوم.
	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	// لنتعرف الآن على فك ترميز بيانات JSON إلى قيم Go. إليك مثالًا
	// على هيكل بيانات عام.
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	// علينا توفير متغير تضع فيه حزمة JSON البيانات الناتجة عن فك
	// الترميز. ستحتوي `map[string]interface{}` هذه على خريطة تربط
	// السلاسل النصية بقيم من أي نوع بيانات.
	var dat map[string]interface{}

	// هذا هو فك الترميز الفعلي، مع التحقق من الأخطاء المرتبطة به.
	// نتجاهل الأخطاء في هذه الأمثلة توخيًا للاختصار، لكن ينبغي في
	// الكود الفعلي التحقق دائمًا من الأخطاء والتعامل معها.
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	// لاستخدام القيم الموجودة في الخريطة الناتجة عن فك الترميز،
	// علينا تحويلها إلى أنواعها المناسبة. نحوّل هنا مثلًا القيمة في
	// `num` إلى النوع المتوقع `float64`.
	num := dat["num"].(float64)
	fmt.Println(num)

	// يتطلب الوصول إلى البيانات المتداخلة سلسلة من التحويلات.
	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	// يمكننا أيضًا فك ترميز JSON إلى أنواع بيانات مخصصة. يضيف ذلك
	// مزيدًا من أمان الأنواع إلى برامجنا، ويلغي الحاجة إلى تأكيدات
	// الأنواع عند الوصول إلى البيانات الناتجة عن فك الترميز.
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	_ = json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	// استخدمنا في الأمثلة أعلاه دائمًا البايتات والسلاسل النصية
	// كقيم وسيطة بين البيانات وتمثيل JSON على الخرج القياسي. يمكننا
	// أيضًا إرسال ترميزات JSON مباشرة إلى كائنات `os.Writer` مثل
	// `os.Stdout`، أو حتى إلى أجسام استجابات HTTP.
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	_ = enc.Encode(d)

	// تُجرى القراءة المتدفقة من كائنات `os.Reader` مثل `os.Stdin`
	// أو أجسام طلبات HTTP باستخدام `json.Decoder`.
	dec := json.NewDecoder(strings.NewReader(str))
	res1 := response2{}
	_ = dec.Decode(&res1)
	fmt.Println(res1)
}
