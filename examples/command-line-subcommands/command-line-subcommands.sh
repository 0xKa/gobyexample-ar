$ go build command-line-subcommands.go 

# استدعِ أولًا الأمر الفرعي `foo`.
$ ./command-line-subcommands foo -enable -name=joe a1 a2
subcommand 'foo'
  enable: true
  name: joe
  tail: [a1 a2]

# جرّب الآن `bar`.
$ ./command-line-subcommands bar -level 8 a1
subcommand 'bar'
  level: 8
  tail: [a1]

# لكن `bar` لن يقبل خيارات `foo`.
$ ./command-line-subcommands bar -enable a1
flag provided but not defined: -enable
Usage of bar:
  -level int
    	level

# سنتعرف تاليًا على متغيرات البيئة، وهي طريقة شائعة أخرى لتحديد
# مُعامِلات البرامج.
