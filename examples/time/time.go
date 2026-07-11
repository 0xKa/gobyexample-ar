// توفر Go دعمًا واسعًا للأوقات والمدد الزمنية. إليك بعض الأمثلة.

package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	// سنبدأ بالحصول على الوقت الحالي.
	now := time.Now()
	p(now)

	// يمكنك إنشاء هيكل `time` بتوفير السنة والشهر واليوم وما إلى
	// ذلك. يرتبط الوقت دائمًا بقيمة `Location`، أي منطقة زمنية.
	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)

	// يمكنك استخراج المكونات المختلفة لقيمة الوقت كما هو متوقع.
	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	// تتوفر أيضًا قيمة `Weekday` التي تمتد من الاثنين إلى الأحد.
	p(then.Weekday())

	// تقارن هذه الأساليب وقتين، وتختبر على الترتيب ما إذا كان الأول
	// يقع قبل الثاني أو بعده أو في الوقت نفسه.
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	// يعيد الأسلوب `Sub` قيمة `Duration` تمثل الفاصل بين وقتين.
	diff := now.Sub(then)
	p(diff)

	// يمكننا حساب طول المدة الزمنية بوحدات مختلفة.
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	// يمكنك استخدام `Add` لتقديم وقت بمقدار مدة معينة، أو استخدامه
	// مع `-` لإرجاع الوقت إلى الوراء بمقدار مدة.
	p(then.Add(diff))
	p(then.Add(-diff))
}
