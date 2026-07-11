// لبعض أدوات سطر الأوامر، مثل `go` أو `git`، عدة *أوامر فرعية*،
// لكل منها مجموعة خيارات خاصة به. يُعد `go build` و`go get` مثلًا
// أمرين فرعيين مختلفين للأداة `go`. تتيح لنا الحزمة `flag` تعريف
// أوامر فرعية بسيطة ذات خيارات خاصة بسهولة.

package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	// نصرّح بأمر فرعي باستخدام الدالة `NewFlagSet`، ثم نعرّف خيارات
	// جديدة خاصة بهذا الأمر الفرعي.
	fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
	fooEnable := fooCmd.Bool("enable", false, "enable")
	fooName := fooCmd.String("name", "", "name")

	// يمكننا تعريف خيارات مدعومة مختلفة لأمر فرعي آخر.
	barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
	barLevel := barCmd.Int("level", 0, "level")

	// يُتوقع أن يكون الأمر الفرعي أول وسيط للبرنامج.
	if len(os.Args) < 2 {
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}

	// تحقق من الأمر الفرعي المستدعى.
	switch os.Args[1] {

	// لكل أمر فرعي، نحلل خياراته ونصل إلى الوسائط الموضعية اللاحقة.
	case "foo":
		fooCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'foo'")
		fmt.Println("  enable:", *fooEnable)
		fmt.Println("  name:", *fooName)
		fmt.Println("  tail:", fooCmd.Args())
	case "bar":
		barCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'bar'")
		fmt.Println("  level:", *barLevel)
		fmt.Println("  tail:", barCmd.Args())
	default:
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}
}
