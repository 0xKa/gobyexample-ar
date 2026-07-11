# استخدم هذه الأوامر لتشغيل المثال.
# (ملاحظة: بسبب قيود Go Playground، لا يمكن تشغيل هذا المثال
# إلا على جهازك المحلي.)
$ mkdir -p folder
$ echo "hello go" > folder/single_file.txt
$ echo "123" > folder/file1.hash
$ echo "456" > folder/file2.hash

$ go run embed-directive.go
hello go
hello go
123
456

