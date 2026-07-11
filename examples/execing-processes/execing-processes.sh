# عند تشغيل برنامجنا، يحل محله الأمر `ls`.
$ go run execing-processes.go
total 16
drwxr-xr-x  4 mark 136B Oct 3 16:29 .
drwxr-xr-x 91 mark 3.0K Oct 3 12:50 ..
-rw-r--r--  1 mark 1.3K Oct 3 16:28 execing-processes.go

# لاحظ أن Go لا توفر دالة `fork` التقليدية في Unix. لا يمثل ذلك
# مشكلة عادة، لأن بدء روتينات Go وتشغيل العمليات واستبدالها يغطي
# معظم حالات استخدام `fork`.
