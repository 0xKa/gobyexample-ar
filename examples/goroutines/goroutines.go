// _روتين Go_ (`goroutine`) هو [خيط تنفيذ](https://ar.wikipedia.org/wiki/%D8%AA%D8%B4%D8%B9%D8%A8_%28%D8%AD%D9%88%D8%B3%D8%A8%D8%A9%29) خفيف الوزن.

package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := range 3 {
		fmt.Println(from, ":", i)
	}
}

func main() {

	// لنفترض أن لدينا استدعاء الدالة `f(s)`. هكذا نستدعيها
	// بالطريقة المعتادة، فتعمل بصورة متزامنة.
	f("direct")

	// لاستدعاء هذه الدالة داخل روتين Go، استخدم `go f(s)`.
	// سيعمل روتين Go الجديد بالتزامن مع الروتين الذي استدعاه.
	go f("goroutine")

	// يمكنك أيضًا بدء روتين Go لاستدعاء دالة مجهولة.
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// يعمل استدعاءا الدالة الآن بصورة غير متزامنة في روتيني
	// Go منفصلين. انتظر حتى ينتهيا (ولطريقة أكثر متانة، استخدم
	// [مجموعة انتظار](waitgroups)).
	time.Sleep(time.Second)
	fmt.Println("done")
}
