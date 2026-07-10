// يمكننا استخدام *مجموعة انتظار* لانتظار انتهاء عدة روتينات Go.

package main

import (
	"fmt"
	"sync"
	"time"
)

// هذه هي الدالة التي سنشغّلها في كل روتين Go.
func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)

	// انتظر لمحاكاة مهمة مكلفة.
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {

	// تُستخدم `WaitGroup` هذه لانتظار انتهاء جميع روتينات Go
	// التي بدأناها هنا. لاحظ أنه عند تمرير `WaitGroup` صراحة إلى
	// الدوال، ينبغي تمريرها *بالمؤشر*.
	var wg sync.WaitGroup

	// ابدأ عدة روتينات Go باستخدام `WaitGroup.Go`.
	for i := 1; i <= 5; i++ {
		wg.Go(func() {
			worker(i)
		})
	}

	// احجب التنفيذ حتى تنتهي جميع روتينات Go التي بدأتها `wg`.
	// ينتهي روتين Go عندما تعود الدالة التي يستدعيها.
	wg.Wait()

	// لاحظ أن هذا الأسلوب لا يوفر طريقة مباشرة لتمرير الأخطاء من
	// العمال. لحالات الاستخدام الأكثر تقدمًا، فكّر في استخدام
	// [الحزمة errgroup](https://pkg.go.dev/golang.org/x/sync/errgroup).
}
