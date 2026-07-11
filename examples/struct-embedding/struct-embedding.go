// تدعم Go _تضمين_ الهياكل والواجهات للتعبير عن _تركيب_
// أكثر سلاسة للأنواع. لا ينبغي الخلط بين ذلك وبين
// [`//go:embed`](embed-directive)، وهو توجيه في Go أُضيف
// في الإصدار 1.16 لتضمين الملفات والمجلدات في الملف
// التنفيذي للتطبيق.

package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

// _يضمّن_ `container` النوع `base`. ويبدو التضمين كحقل
// بلا اسم.
type container struct {
	base
	str string
}

func main() {

	// عند إنشاء الهياكل باستخدام القيم الحرفية، علينا تهيئة
	// التضمين صراحةً؛ ويؤدي النوع المضمّن هنا دور اسم الحقل.
	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}

	// يمكننا الوصول إلى حقول `base` مباشرةً عبر `co`،
	// مثل `co.num`.
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	// بدلًا من ذلك، يمكننا كتابة المسار كاملًا باستخدام
	// اسم النوع المضمّن.
	fmt.Println("also num:", co.base.num)

	// بما أن `container` يضمّن `base`، تصبح أساليب `base`
	// أيضًا أساليب لـ`container`. نستدعي هنا مباشرةً عبر
	// `co` أسلوبًا مضمّنًا من `base`.
	fmt.Println("describe:", co.describe())

	type describer interface {
		describe() string
	}

	// يمكن استخدام تضمين الهياكل ذات الأساليب لجعل هياكل
	// أخرى تطبق واجهات. نرى هنا أن `container` يطبق الآن
	// واجهة `describer` لأنه يضمّن `base`.
	var d describer = co
	fmt.Println("describer:", d.describe())
}
