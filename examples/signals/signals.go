// نريد أحيانًا أن تتعامل برامج Go بذكاء مع [إشارات
// Unix](https://en.wikipedia.org/wiki/Unix_signal). فقد نريد مثلًا
// إيقاف خادم بسلاسة عند تلقيه `SIGTERM`، أو إيقاف أداة سطر أوامر
// عن معالجة المدخلات إذا تلقت `SIGINT`. إليك طريقة حديثة لمعالجة
// الإشارات باستخدام السياقات.

package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"
)

func main() {
	// تعيد `signal.NotifyContext` سياقًا يُلغى عند وصول إحدى الإشارات
	// المدرجة.
	ctx, stop := signal.NotifyContext(
		context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// سينتظر البرنامج هنا حتى يستقبل إحدى الإشارات المحددة.
	fmt.Println("awaiting signal")
	<-ctx.Done()

	// تعيد `context.Cause` سبب إلغاء السياق. وعندما يكون الإلغاء
	// ناتجًا عن إشارة، تتضمن النتيجة قيمة الإشارة.
	fmt.Println()
	fmt.Println(context.Cause(ctx))
	fmt.Println("exiting")
}
