# لاحظ أن الشرائح، رغم اختلافها عن المصفوفات في النوع، تُعرض
# بطريقة مشابهة عند استخدام `fmt.Println`.
$ go run slices.go
uninit: [] true true
emp: [  ] len: 3 cap: 3
set: [a b c]
get: c
len: 3
apd: [a b c d e f]
cpy: [a b c d e f]
sl1: [c d e]
sl2: [a b c d e]
sl3: [c d e f]
dcl: [g h i]
t == t2
2d:  [[0] [1 2] [2 3 4]]

# اقرأ هذه [التدوينة المميزة](https://go.dev/blog/slices-intro)
# من فريق Go لمزيد من التفاصيل حول تصميم الشرائح
# وتنفيذها في Go.

# بعد أن تعرفنا على المصفوفات والشرائح، سننتقل إلى هيكل
# بيانات مدمجة أساسية أخرى في Go، وهي الخرائط.
