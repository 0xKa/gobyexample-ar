# لتشغيل البرنامج، ضع الكود في `hello-world.go` ثم
# استخدم `go run`.
$ go run hello-world.go
hello world

# نحتاج أحيانًا إلى بناء برامجنا على شكل ملفات
# تنفيذية. يمكننا فعل ذلك باستخدام `go build`.
$ go build hello-world.go
$ ls
hello-world	hello-world.go

# يمكننا بعد ذلك تشغيل الملف التنفيذي الناتج مباشرة.
$ ./hello-world
hello world

# بعدما أصبح بإمكاننا تشغيل برامج Go الأساسية وبنائها،
# فلنتعرف أكثر على اللغة.
