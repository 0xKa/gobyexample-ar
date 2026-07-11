// توفر عناوين URL [طريقة موحدة لتحديد مواقع الموارد](https://adam.herokuapp.com/past/2010/3/30/urls_are_the_uniform_way_to_locate_resources/).
// هكذا تحلل عناوين URL في Go.

package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {

	// سنحلل عنوان URL هذا، وهو يتضمن مخططًا ومعلومات مصادقة
	// ومضيفًا ومنفذًا ومسارًا ومُعامِلات استعلام وجزءًا.
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	// حلل عنوان URL وتأكد من عدم وجود أخطاء.
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	// الوصول إلى المخطط مباشر.
	fmt.Println(u.Scheme)

	// يحتوي `User` على جميع معلومات المصادقة. استدعِ عليه
	// `Username` و`Password` للحصول على القيم المنفردة.
	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	p, _ := u.User.Password()
	fmt.Println(p)

	// يحتوي `Host` على اسم المضيف والمنفذ إن وُجد. استخدم
	// `SplitHostPort` لاستخراجهما.
	fmt.Println(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host)
	fmt.Println(port)

	// نستخرج هنا `path` والجزء الواقع بعد `#`.
	fmt.Println(u.Path)
	fmt.Println(u.Fragment)

	// للحصول على مُعامِلات الاستعلام في سلسلة نصية بالصيغة `k=v`،
	// استخدم `RawQuery`. يمكنك أيضًا تحليل مُعامِلات الاستعلام إلى
	// خريطة. تربط خرائط الاستعلام المحللة السلاسل النصية بشرائح
	// من السلاسل النصية، لذا استخدم الفهرس `[0]` إذا أردت القيمة
	// الأولى فقط.
	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0])
}
