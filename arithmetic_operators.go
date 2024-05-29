// package main //func main มีได้แค่ตัวเดียวใน 1 project

// import "fmt"

// func main() {
// 	//13. ตัวดำเนินการทางคณิตศาสตร์
// 	//13.1 ถ้าเป็นค่าคงที่ ไม่สามารถเป็นชนิด float ได้
// 	//13.2 การบวก-ลบ-คูณ-หาร ต้องเป็นชนิดเดียวกันถึงจะอนุญาตให้ดำเนินการได้
// 	//เช่น int+int,float32-float32,float64*float64
// 	//13.3 ถ้าชนิดเป็น int ผลลัพธ์จะเป็น int
// 	//ถ้าชนิดเป็น float32 ผลลัพธ์ใดที่มีเลขหลังจุดทศนิยมจะได้เลขหลังจุดทศนิยม 8 ตำแหน่ง
// 	//ถ้าชนิดเป็น float64 ผลลัพธ์ใดที่มีเลขหลังจุดทศนิยมจะได้เลขหลังจุดทศนิยม 16 ตำแหน่ง
// 	//13.4 การหารเอาเศษต้องเป็น int เท่านั้น

// 	const number1 int = 25
// 	const number2 int = 60
// 	fmt.Println("ผลบวก =", number1+number2)
// 	fmt.Printf("Type ของผลบวก %T \n\n", number1+number2)

// 	fmt.Println("ผลลบ =", number1-number2)
// 	fmt.Printf("Type ของผลลบ %T \n\n", number1-number2)

// 	fmt.Println("ผลคูณ =", number1*number2)
// 	fmt.Printf("Type ของผลคูณ %T \n\n", number1*number2)

// 	//if number1 and number2 or one
// 	fmt.Println("ผลหาร =", number1/number2)
// 	fmt.Printf("Type ของผลหาร %T \n\n", number1/number2)

// 	//Only int
// 	fmt.Println("ผลการหารเอาเศษ =", number1%number2)
// 	fmt.Printf("Type ของผลการหารเอาเศษ %T \n\n", number1%number2)

// }
