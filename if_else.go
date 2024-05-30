package main

import "fmt"

func main() {
	var score int
	fmt.Printf("โปรดกรอบคะแนนของคุณ : ")
	fmt.Scanf("%d", &score)
	if score >= 50 {
		fmt.Println("คะแนนคุณคือ : ", score)
		fmt.Println("ยินดีด้วยคุณสอบผ่าน")
	} else {
		fmt.Println("คะแนนของคุณคือ :", score)
		fmt.Println("พยายามใหม่นะ คุณสอบไม่ผ่าน")
	}
}
