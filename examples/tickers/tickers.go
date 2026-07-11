// تُستخدم [المؤقتات](timers) عندما تريد تنفيذ شيء مرة واحدة
// مستقبلًا، بينما تُستخدم _المؤقتات الدورية_ عندما تريد تنفيذه
// مرارًا على فواصل زمنية منتظمة. إليك مثالًا على مؤقت دوري
// يرسل نبضات بانتظام حتى نوقفه.

package main

import (
	"fmt"
	"time"
)

func main() {

	// تستخدم المؤقتات الدورية آلية شبيهة بالمؤقتات: قناة تُرسل
	// إليها القيم. سنستخدم هنا عبارة `select` المدمجة على القناة
	// لانتظار القيم عند وصولها كل 500 مللي ثانية.
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	// يمكن إيقاف المؤقتات الدورية مثل المؤقتات. بعد إيقاف المؤقت
	// الدوري، لن تتلقى قناته أي قيم أخرى. سنوقف مؤقتنا بعد 1600
	// مللي ثانية.
	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
