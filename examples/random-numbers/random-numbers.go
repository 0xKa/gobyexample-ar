// توفر الحزمة `math/rand/v2` في Go توليد
// [الأعداد شبه العشوائية](https://en.wikipedia.org/wiki/Pseudorandom_number_generator).

package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {

	// تعيد `rand.IntN` مثلًا عددًا عشوائيًا `n` من النوع `int`،
	// بحيث `0 <= n < 100`.
	fmt.Print(rand.IntN(100), ",")
	fmt.Print(rand.IntN(100))
	fmt.Println()

	// تعيد `rand.Float64` قيمة عشوائية `f` من النوع `float64`، بحيث
	// `0.0 <= f < 1.0`.
	fmt.Println(rand.Float64())

	// يمكن استخدام ذلك لتوليد أعداد فاصلة عائمة عشوائية في نطاقات
	// أخرى، مثل `5.0 <= f' < 10.0`.
	fmt.Print((rand.Float64()*5)+5, ",")
	fmt.Print((rand.Float64() * 5) + 5)
	fmt.Println()

	// إذا أردت بذرة معلومة، فأنشئ `rand.Source` جديدًا ومرره إلى
	// دالة الإنشاء `New`. تنشئ `NewPCG` مصدر
	// [PCG](https://en.wikipedia.org/wiki/Permuted_congruential_generator)
	// جديدًا يتطلب بذرة مكونة من عددين من النوع `uint64`.
	s2 := rand.NewPCG(42, 1024)
	r2 := rand.New(s2)
	fmt.Print(r2.IntN(100), ",")
	fmt.Print(r2.IntN(100))
	fmt.Println()

	s3 := rand.NewPCG(42, 1024)
	r3 := rand.New(s3)
	fmt.Print(r3.IntN(100), ",")
	fmt.Print(r3.IntN(100))
	fmt.Println()
}
