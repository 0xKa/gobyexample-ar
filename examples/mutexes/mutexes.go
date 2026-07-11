// رأينا في المثال السابق كيفية إدارة حالة عداد بسيطة باستخدام
// [العمليات الذرية](atomic-counters). للحالات الأكثر تعقيدًا،
// يمكننا استخدام [قفل استبعاد متبادل](https://ar.wikipedia.org/wiki/استبعاد_التشارك)
// للوصول إلى البيانات بأمان من عدة روتينات Go.

package main

import (
	"fmt"
	"sync"
)

// يحتوي `Container` على خريطة عدادات. ولأننا نريد تحديثها
// بالتزامن من عدة روتينات Go، نضيف `Mutex` لمزامنة الوصول.
// لاحظ أنه يجب عدم نسخ أقفال الاستبعاد المتبادل، ولذلك ينبغي
// تمرير هذا الهيكل بالمؤشر إذا مُرّر بين الدوال.
type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *Container) inc(name string) {
	// اقفل قفل الاستبعاد المتبادل قبل الوصول إلى `counters`،
	// وافتحه في نهاية الدالة باستخدام عبارة [`defer`](defer).
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

func main() {
	c := Container{
		// لاحظ أن القيمة الصفرية لقفل الاستبعاد المتبادل قابلة
		// للاستخدام كما هي، فلا حاجة إلى تهيئته هنا.
		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup

	// تزيد هذه الدالة عدادًا محدد الاسم داخل حلقة تكرار.
	doIncrement := func(name string, n int) {
		for range n {
			c.inc(name)
		}
	}

	// شغّل عدة روتينات Go بالتزامن. لاحظ أنها جميعًا تصل إلى
	// `Container` نفسه، وأن اثنين منها يصلان إلى العداد نفسه.
	wg.Go(func() {
		doIncrement("a", 10000)
	})

	wg.Go(func() {
		doIncrement("a", 10000)
	})

	wg.Go(func() {
		doIncrement("b", 10000)
	})

	// انتظر انتهاء روتينات Go.
	wg.Wait()
	fmt.Println(c.counters)
}
