// عمليات الإرسال والاستقبال الأساسية على القنوات حاجزة. لكن
// يمكننا استخدام `select` مع فرع `default` لتنفيذ عمليات إرسال
// واستقبال _غير حاجزة_، بل وحتى عبارات `select` متعددة الفروع
// وغير حاجزة.

package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	// هذه عملية استقبال غير حاجزة. إذا توفرت قيمة في `messages`،
	// فستختار `select` فرع `<-messages` مع تلك القيمة. وإلا
	// فستختار فرع `default` فورًا.
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	// تعمل عملية الإرسال غير الحاجزة بطريقة مماثلة. لا يمكن هنا
	// إرسال `msg` إلى القناة `messages`، لأن القناة بلا مخزن
	// مؤقت ولا يوجد مستقبِل. لذلك يُختار فرع `default`.
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	// يمكننا استخدام عدة فروع `case` قبل فرع `default` لتنفيذ
	// `select` متعددة الفروع وغير حاجزة. نحاول هنا إجراء عمليتي
	// استقبال غير حاجزتين من `messages` و`signals`.
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
