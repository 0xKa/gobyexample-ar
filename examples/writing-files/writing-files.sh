# جرّب تشغيل كود كتابة الملفات.
$ go run writing-files.go 
wrote 5 bytes
wrote 7 bytes
wrote 9 bytes

# ثم تحقق من محتويات الملفات المكتوبة.
$ cat /tmp/dat1
hello
go
$ cat /tmp/dat2
some
writes
buffered

# سنتعرف تاليًا على تطبيق بعض أفكار إدخال الملفات وإخراجها التي
# رأيناها للتو على تدفقي `stdin` و`stdout`.
