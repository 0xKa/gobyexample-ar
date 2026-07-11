# يوضح تشغيل البرنامج أننا نحصل على قيمة `FOO` التي عيّناها في
# البرنامج، بينما تكون `BAR` فارغة.
$ go run environment-variables.go
FOO: 1
BAR: 

# تعتمد قائمة المفاتيح في البيئة على جهازك.
TERM_PROGRAM
PATH
SHELL
...
FOO

# إذا عيّنا `BAR` في البيئة أولًا، فسيحصل البرنامج عند تشغيله على
# تلك القيمة.
$ BAR=2 go run environment-variables.go
FOO: 1
BAR: 2
...
