# لتجربة مرشح الأسطر، أنشئ أولًا ملفًا يحتوي على بضعة أسطر بأحرف
# إنجليزية صغيرة.
$ echo 'hello'   > /tmp/lines
$ echo 'filter' >> /tmp/lines

# ثم استخدم مرشح الأسطر للحصول على أسطر بأحرف إنجليزية كبيرة.
$ cat /tmp/lines | go run line-filters.go
HELLO
FILTER
