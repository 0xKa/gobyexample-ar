// _الأنواع المعدّدة_ (التعدادات) حالة خاصة من
// [الأنواع المجموعية](https://en.wikipedia.org/wiki/Algebraic_data_type).
// التعداد نوع له عدد ثابت من القيم الممكنة، ولكل منها اسم
// مميز. لا تملك Go نوع تعداد كميزة لغوية مستقلة، لكن يسهل
// تطبيق التعدادات باستخدام أساليب اللغة المتعارف عليها.

package main

import "fmt"

// النوع الأساسي لتعدادنا `ServerState` هو `int`.
type ServerState int

// تُعرّف القيم الممكنة لـ`ServerState` كثوابت. تولد الكلمة
// الخاصة [iota](https://go.dev/ref/spec#Iota) قيمًا ثابتة
// متتالية تلقائيًا؛ وهي في هذه الحالة `0` و`1` و`2`
// وهكذا.
const (
	StateIdle ServerState = iota
	StateConnected
	StateError
	StateRetrying
)

// بتطبيق واجهة [fmt.Stringer](https://pkg.go.dev/fmt#Stringer)،
// يمكن طباعة قيم `ServerState` أو تحويلها إلى سلاسل نصية.
//
// قد يصبح هذا مرهقًا عند وجود قيم ممكنة كثيرة. في هذه
// الحالات، يمكن استخدام [أداة stringer](https://pkg.go.dev/golang.org/x/tools/cmd/stringer)
// مع `go:generate` لأتمتة العملية. راجع
// [هذه المقالة](https://eli.thegreenplace.net/2021/a-comprehensive-guide-to-go-generate)
// لشرح أطول.
var stateName = map[ServerState]string{
	StateIdle:      "idle",
	StateConnected: "connected",
	StateError:     "error",
	StateRetrying:  "retrying",
}

func (ss ServerState) String() string {
	return stateName[ss]
}

func main() {
	ns := transition(StateIdle)
	fmt.Println(ns)
	// إذا كانت لدينا قيمة من النوع `int`، فلا يمكننا تمريرها
	// إلى `transition`، إذ سيبلغ المترجم عن عدم تطابق النوع.
	// يوفر ذلك قدرًا من سلامة أنواع التعدادات وقت الترجمة.

	ns2 := transition(ns)
	fmt.Println(ns2)
}

// تحاكي `transition` انتقال حالة لخادم؛ فتستقبل الحالة
// الحالية وتعيد حالة جديدة.
func transition(s ServerState) ServerState {
	switch s {
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:
		// افترض أننا نتحقق هنا من بعض الشروط لتحديد
		// الحالة التالية...
		return StateIdle
	case StateError:
		return StateError
	default:
		panic(fmt.Errorf("unknown state: %s", s))
	}
}
