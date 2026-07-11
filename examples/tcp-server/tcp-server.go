// توفر الحزمة `net` الأدوات التي نحتاجها لبناء خوادم مقابس TCP
// بسهولة.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {

	// تبدأ `net.Listen` الخادم على الشبكة المحددة (TCP) والعنوان
	// المحدد (المنفذ 8090 على جميع واجهات الشبكة).
	listener, err := net.Listen("tcp", ":8090")
	if err != nil {
		log.Fatal("Error listening:", err)
	}

	// أغلق المستمع لتحرير المنفذ عند خروج التطبيق.
	defer listener.Close()

	// كرر إلى أجل غير مسمى لقبول اتصالات عملاء جديدة.
	for {
		// انتظر اتصالًا.
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting conn:", err)
			continue
		}

		// نستخدم هنا روتين Go لمعالجة الاتصال، حتى تتمكن الحلقة
		// الرئيسية من مواصلة قبول المزيد من الاتصالات.
		go handleConnection(conn)
	}
}

// تعالج `handleConnection` اتصالًا واحدًا لعميل، فتقرأ سطرًا
// نصيًا واحدًا من العميل وتعيد استجابة.
func handleConnection(conn net.Conn) {
	// يحرر إغلاق الاتصال الموارد عند انتهاء تفاعلنا مع العميل.
	defer conn.Close()

	// استخدم `bufio.NewReader` لقراءة سطر بيانات واحد من العميل،
	// ينتهي بمحرف سطر جديد.
	reader := bufio.NewReader(conn)
	message, err := reader.ReadString('\n')
	if err != nil {
		log.Printf("Read error: %v", err)
		return
	}

	// أنشئ استجابة وأرسلها إلى العميل، موضحًا الاتصال ثنائي الاتجاه.
	ackMsg := strings.ToUpper(strings.TrimSpace(message))
	response := fmt.Sprintf("ACK: %s\n", ackMsg)
	_, err = conn.Write([]byte(response))
	if err != nil {
		log.Printf("Server write error: %v", err)
	}
}
