// يُعد [تحديد المعدل](https://en.wikipedia.org/wiki/Rate_limiting)
// آلية مهمة للتحكم في استخدام الموارد والحفاظ على جودة الخدمة.
// تدعم Go تحديد المعدل بأناقة باستخدام روتينات Go والقنوات
// و[المؤقتات الدورية](tickers).

package main

import (
	"fmt"
	"time"
)

func main() {

	// سنتعرف أولًا على تحديد المعدل الأساسي. لنفترض أننا نريد
	// تقييد معالجتنا للطلبات الواردة. سنخدم هذه الطلبات من قناة
	// تحمل الاسم نفسه.
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// ستستقبل القناة `limiter` قيمة كل 200 مللي ثانية. وهي المنظّم
	// في آلية تحديد المعدل لدينا.
	limiter := time.Tick(200 * time.Millisecond)

	// بحجب التنفيذ عند الاستقبال من القناة `limiter` قبل خدمة كل
	// طلب، نقيد أنفسنا بطلب واحد كل 200 مللي ثانية.
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// قد نرغب في السماح بدفعات قصيرة من الطلبات ضمن آلية تحديد
	// المعدل، مع الحفاظ على المعدل الإجمالي. يمكننا تحقيق ذلك
	// بتخزين قناة تحديد المعدل مؤقتًا. ستسمح القناة
	// `burstyLimiter` بدفعات تصل إلى 3 أحداث.
	burstyLimiter := make(chan time.Time, 3)

	// املأ القناة لتمثيل الدفعات المسموح بها.
	for range 3 {
		burstyLimiter <- time.Now()
	}

	// سنحاول كل 200 مللي ثانية إضافة قيمة جديدة إلى
	// `burstyLimiter`، حتى حدها البالغ 3 قيم.
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	// حاكِ الآن 5 طلبات واردة أخرى. ستستفيد الطلبات الثلاثة
	// الأولى من قدرة `burstyLimiter` على معالجة الدفعات.
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
