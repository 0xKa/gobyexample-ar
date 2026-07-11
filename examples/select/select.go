// تتيح لك عبارة _`select`_ في Go انتظار عدة عمليات على
// القنوات. يُعد الجمع بين روتينات Go والقنوات و`select` من
// مزايا Go القوية.

package main

import (
	"fmt"
	"time"
)

func main() {

	// سنختار في مثالنا بين قناتين.
	c1 := make(chan string)
	c2 := make(chan string)

	// ستستقبل كل قناة قيمة بعد مدة معينة، لمحاكاة عمليات RPC
	// حاجزة مثلًا تعمل داخل روتينات Go متزامنة.
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	// سنستخدم `select` لانتظار القيمتين معًا، وطباعة كل منهما
	// عند وصولها.
	for range 2 {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}
