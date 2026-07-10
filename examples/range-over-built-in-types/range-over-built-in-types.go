// يجتاز `range` عناصر مجموعة متنوعة من هياكل البيانات
// المدمجة. لنرَ كيفية استخدام `range` مع بعض هياكل
// البيانات التي تعرفنا عليها.

package main

import "fmt"

func main() {

	// نستخدم هنا `range` لجمع الأرقام في شريحة. تعمل
	// المصفوفات بالطريقة نفسها أيضًا.
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// يوفر `range` عند استخدامه مع المصفوفات والشرائح فهرس
	// كل عنصر وقيمته. لم نحتج إلى الفهرس أعلاه، فتجاهلناه
	// باستخدام المعرّف الفارغ `_`. لكننا نحتاج إلى الفهارس
	// أحيانًا.
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// يجتاز `range` أزواج المفاتيح والقيم في الخريطة.
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// يمكن لـ`range` أيضًا اجتياز مفاتيح الخريطة وحدها.
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// يجتاز `range` نقاط ترميز Unicode في السلاسل النصية.
	// القيمة الأولى هي فهرس البايت الذي يبدأ عنده محرف
	// `rune`، والثانية هي محرف `rune` نفسه. راجع
	// [السلاسل النصية والمحارف](strings-and-runes) لمزيد
	// من التفاصيل.
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
