# عند تشغيل هذا البرنامج، نرى أولًا خرج الاستدعاء الحاجز،
# ثم خرج روتيني Go. قد تتداخل أسطر خرجهما لأن بيئة تشغيل Go
# تشغّلهما بالتزامن.
$ go run goroutines.go
direct : 0
direct : 1
direct : 2
goroutine : 0
going
goroutine : 1
goroutine : 2
done

# سنتعرف تاليًا على ما يكمل روتينات Go في برامج Go
# المتزامنة: القنوات.
