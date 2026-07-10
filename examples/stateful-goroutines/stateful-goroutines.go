// استخدمنا في المثال السابق القفل الصريح باستخدام
// [أقفال الاستبعاد المتبادل](mutexes) لمزامنة وصول عدة روتينات
// Go إلى حالة مشتركة. يمكن بدلًا من ذلك استخدام مزايا المزامنة
// المدمجة في روتينات Go والقنوات لتحقيق النتيجة نفسها. يتوافق
// هذا الأسلوب القائم على القنوات مع نهج Go في مشاركة الذاكرة
// عن طريق الاتصال، وامتلاك روتين Go واحد فقط لكل جزء من البيانات.

package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

// سيمتلك روتين Go واحد حالتنا في هذا المثال، ما يضمن عدم تلف
// البيانات بسبب الوصول المتزامن. لقراءة هذه الحالة أو كتابتها،
// سترسل روتينات Go الأخرى رسائل إلى الروتين المالك وتتلقى
// الردود المقابلة. يغلّف الهيكلان `readOp` و`writeOp` هذه
// الطلبات، ويوفران وسيلة للروتين المالك كي يرد عليها.
type readOp struct {
	key  int
	resp chan int
}
type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {

	// سنعد، كما سبق، عدد العمليات التي ننفذها.
	var readOps uint64
	var writeOps uint64

	// ستستخدم روتينات Go الأخرى القناتين `reads` و`writes`
	// لإصدار طلبات القراءة والكتابة على الترتيب.
	reads := make(chan readOp)
	writes := make(chan writeOp)

	// هذا هو روتين Go الذي يمتلك `state`، وهي خريطة كما في المثال
	// السابق، لكنها الآن خاصة بروتين Go ذي الحالة. يختار هذا
	// الروتين مرارًا بين القناتين `reads` و`writes`، ويرد على
	// الطلبات عند وصولها. لتنفيذ الرد، ينفذ أولًا العملية المطلوبة،
	// ثم يرسل قيمة عبر قناة الرد `resp` للإشارة إلى النجاح (ومعها
	// القيمة المطلوبة في حالة `reads`).
	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	// يبدأ هذا تشغيل 100 روتين Go لإصدار طلبات قراءة إلى الروتين
	// المالك للحالة عبر القناة `reads`. تتطلب كل قراءة إنشاء `readOp`
	// وإرساله عبر القناة `reads`، ثم استقبال النتيجة عبر القناة
	// `resp` المقدمة.
	for range 100 {
		go func() {
			for {
				read := readOp{
					key:  rand.Intn(5),
					resp: make(chan int)}
				reads <- read
				<-read.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// نبدأ أيضًا 10 روتينات للكتابة باستخدام أسلوب مماثل.
	for range 10 {
		go func() {
			for {
				write := writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool)}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// دع روتينات Go تعمل لمدة ثانية.
	time.Sleep(time.Second)

	// احفظ أخيرًا أعداد العمليات واعرضها.
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)
}
