# شغّل خادم TCP في الخلفية.
$ go run tcp-server.go &

# أرسل البيانات والتقط الاستجابة باستخدام netcat.
$ echo "Hello from netcat" | nc localhost 8090
ACK: HELLO FROM NETCAT

