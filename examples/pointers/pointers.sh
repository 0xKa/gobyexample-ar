# لا تغير `zeroval` قيمة `i` في `main`، بينما تغيرها
# `zeroptr` لأنها تملك مرجعًا إلى عنوان ذاكرة ذلك المتغير.
$ go run pointers.go
initial: 1
zeroval: 1
zeroptr: 0
pointer: 0x42131100
