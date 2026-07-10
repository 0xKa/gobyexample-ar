$ go run channel-synchronization.go      
working...done                  

# إذا حذفت السطر `<- done` من هذا البرنامج، فقد يخرج البرنامج
# قبل أن ينهي `worker` عمله، أو حتى قبل أن يبدأ في بعض الحالات.
