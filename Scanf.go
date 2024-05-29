package main

import "fmt"

func main() {
	//15. การรับค่าจากคีย์บอร์ดด้วย Scanf
	var ans string
	fmt.Print("วันนี้คุณทานข้าวกับอะไร??")
	fmt.Scanf("%s", &ans)

	fmt.Println("ทานข้าวกับ", ans, "หรอ ขอให้วันนี้เป็นวันที่ดีสำหรับคุณ")

}
