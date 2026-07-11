// سنتعرف في هذا المثال على كيفية تنفيذ _مجموعة عمال_ باستخدام
// روتينات Go والقنوات.

package main

import (
	"fmt"
	"time"
)

// هذا هو العامل الذي سنشغّل عدة نسخ متزامنة منه. سيستقبل هؤلاء
// العمال العمل عبر القناة `jobs`، ويرسلون النتائج المقابلة عبر
// `results`. سننتظر ثانية لكل مهمة لمحاكاة مهمة مكلفة.
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {

	// لاستخدام مجموعة العمال، علينا إرسال العمل إليهم وجمع
	// نتائجهم. ننشئ قناتين لهذا الغرض.
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// يبدأ هذا تشغيل 3 عمال، يكونون محجوبين أولًا لعدم وجود مهام بعد.
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// نرسل هنا 5 مهام عبر `jobs`، ثم نغلق تلك القناة باستخدام
	// `close` للإشارة إلى أن هذا كل العمل المتاح لدينا.
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// نجمع أخيرًا جميع نتائج العمل. يضمن ذلك أيضًا انتهاء روتينات
	// Go العاملة. ومن الطرق البديلة لانتظار عدة روتينات Go
	// استخدام [مجموعة انتظار](waitgroups).
	for a := 1; a <= numJobs; a++ {
		<-results
	}
}
