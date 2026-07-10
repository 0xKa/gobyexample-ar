// _الهياكل_ في Go مجموعات من الحقول محددة الأنواع.
// تفيد في تجميع البيانات معًا لتكوين سجلات.

package main

import "fmt"

// يحتوي نوع الهيكل `person` على الحقلين `name` و`age`.
type person struct {
	name string
	age  int
}

// تنشئ `newPerson` هيكل `person` جديدًا بالاسم المعطى.
func newPerson(name string) *person {
	// تستخدم Go جمع المهملات؛ لذلك يمكنك بأمان إعادة مؤشر
	// إلى متغير محلي، إذ لن ينظفه جامع المهملات إلا عندما
	// تنعدم المراجع النشطة إليه.
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {

	// تنشئ هذه الصيغة هيكلًا جديدًا.
	fmt.Println(person{"Bob", 20})

	// يمكنك تسمية الحقول عند تهيئة هيكل.
	fmt.Println(person{name: "Alice", age: 30})

	// تأخذ الحقول غير المذكورة قيمها الصفرية.
	fmt.Println(person{name: "Fred"})

	// تعطي البادئة `&` مؤشرًا إلى الهيكل.
	fmt.Println(&person{name: "Ann", age: 40})

	// من المتعارف عليه حصر إنشاء الهياكل الجديدة في دوال إنشاء.
	fmt.Println(newPerson("Jon"))

	// استخدم النقطة للوصول إلى حقول الهيكل.
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	// يمكنك أيضًا استخدام النقاط مع مؤشرات الهياكل؛ إذ تفك
	// Go إشارة المؤشرات تلقائيًا.
	sp := &s
	fmt.Println(sp.age)

	// الهياكل قابلة للتغيير.
	sp.age = 51
	fmt.Println(sp.age)

	// إذا لم يُستخدم نوع هيكل إلا لقيمة واحدة، فلا يلزم أن
	// نسميه؛ إذ يمكن أن يكون للقيمة نوع هيكل مجهول. يشيع
	// استخدام هذه التقنية في
	// [الاختبارات القائمة على جداول](testing-and-benchmarking).
	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)
}
