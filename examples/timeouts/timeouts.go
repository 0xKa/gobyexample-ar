// _المهل الزمنية_ مهمة للبرامج التي تتصل بموارد خارجية، أو
// التي تحتاج لسبب آخر إلى تقييد وقت التنفيذ. يسهل تنفيذ المهل
// الزمنية في Go بأناقة بفضل القنوات و`select`.

package main

import (
	"fmt"
	"time"
)

func main() {

	// لنفترض في مثالنا أننا ننفذ استدعاءً خارجيًا يعيد نتيجته
	// عبر القناة `c1` بعد ثانيتين. لاحظ أن القناة مخزنة مؤقتًا،
	// ولذلك تكون عملية الإرسال داخل روتين Go غير حاجزة. هذا نمط
	// شائع لمنع تسرب روتين Go إذا لم تُقرأ القناة قط.
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	// تنفّذ `select` هنا مهلة زمنية. تنتظر `res := <-c1`
	// النتيجة، بينما تنتظر `<-time.After` إرسال قيمة بعد مهلة
	// مدتها ثانية واحدة. ولأن `select` تتابع أول عملية استقبال
	// تصبح جاهزة، فسننفذ فرع المهلة إذا استغرقت العملية أكثر من
	// الثانية المسموح بها.
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	// إذا سمحنا بمهلة أطول مدتها 3 ثوانٍ، فستنجح عملية الاستقبال
	// من `c2` وسنطبع النتيجة.
	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}
