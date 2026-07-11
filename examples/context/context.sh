# شغّل الخادم في الخلفية.
$ go run context.go &

# حاكِ طلب عميل إلى `/hello`، واضغط `Ctrl+C` بعد وقت قصير من البدء
# لإرسال إشارة إلغاء.
$ curl localhost:8090/hello
server: hello handler started
^C
server: context canceled
server: hello handler ended
