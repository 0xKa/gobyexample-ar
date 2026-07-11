# عند تشغيل هذا البرنامج، سيُحجب منتظرًا إشارة. يمكننا بالضغط على
# `ctrl-C`، التي تعرضها الطرفية على صورة `^C`، إرسال إشارة
# `SIGINT`، فيطبع البرنامج سبب الإلغاء ثم يخرج.
$ go run signals.go
awaiting signal
^C
interrupt signal received
exiting
