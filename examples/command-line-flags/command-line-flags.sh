# لتجربة برنامج خيارات سطر الأوامر، يُفضل أولًا ترجمته، ثم تشغيل
# الملف التنفيذي الناتج مباشرة.
$ go build command-line-flags.go

# جرّب البرنامج المبني بإعطائه أولًا قيمًا لجميع الخيارات.
$ ./command-line-flags -word=opt -numb=7 -fork -svar=flag
word: opt
numb: 7
fork: true
svar: flag
tail: []

# لاحظ أن الخيارات التي لا تذكرها تأخذ قيمها الافتراضية تلقائيًا.
$ ./command-line-flags -word=opt
word: opt
numb: 42
fork: false
svar: bar
tail: []

# يمكن توفير الوسائط الموضعية اللاحقة بعد أي خيارات.
$ ./command-line-flags -word=opt a1 a2 a3
word: opt
...
tail: [a1 a2 a3]

# لاحظ أن الحزمة `flag` تشترط ظهور جميع الخيارات قبل الوسائط
# الموضعية، وإلا فستُفسر الخيارات على أنها وسائط موضعية.
$ ./command-line-flags -word=opt a1 a2 a3 -numb=7
word: opt
numb: 42
fork: false
svar: bar
tail: [a1 a2 a3 -numb=7]

# استخدم الخيار `-h` أو `--help` للحصول على نص مساعدة مولّد
# تلقائيًا لبرنامج سطر الأوامر.
$ ./command-line-flags -h
Usage of ./command-line-flags:
  -fork=false: a bool
  -numb=42: an int
  -svar="bar": a string var
  -word="foo": a string

# إذا قدمت خيارًا لم يُعرّف في الحزمة `flag`، فسيطبع البرنامج رسالة
# خطأ ويعرض نص المساعدة مجددًا.
$ ./command-line-flags -wat
flag provided but not defined: -wat
Usage of ./command-line-flags:
...
