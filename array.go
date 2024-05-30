package main

import "fmt"

func main() {

	var number [6]int
	var str [3]string
	var fruit [3]string = [3]string{"orange", "apple", "banana"}
	ex := [...]int{10, 20, 30}
	fmt.Println("Array")
	fmt.Println("data in zero value of number :", number)
	number[0] = 100
	fmt.Println("number after add data :", number)
	fmt.Println("data in zero value of str :", str)
	fmt.Println("ex  :", ex)
	fmt.Println("index 0 of fruit :", fruit[0])
	fmt.Println("index 1 of fruit :", fruit[1])
	fmt.Println("index 2 of fruit :", fruit[2])
	fruit[2] = "mango"
	fmt.Println("change index 2 of fruit :", fruit)
	fmt.Println("จำนวนสมาชิกใน fruit :", len(fruit))
	fmt.Println("--------------------------------")

	fmt.Println("Slices")
	var names []string = []string{"Dave", "Alley"}
	fmt.Println("names :", names)
	//เพิ่มสมาชิก
	names = append(names, "John")
	fmt.Println("names after append :", names)
	//เข้าถึงสมาชิก โดย index
	fmt.Println("index 0 of names :", names[0])
	fmt.Println("index 1 of names :", names[1])
	fmt.Println("index 2 of names :", names[2])
	//การเข้าถึงสมาชิกแบบกำหนดช่วง
	fmt.Println("names[:1] = เอาสมาชิกตั้งแต่ 0 ไม่เอา 1 :", names[:1])
	fmt.Println("names[0:2] = เอาสมาชิก 0,1 ไม่เอา 2 :", names[0:2])
	fmt.Println("สามชิก 1-สุดท้าย :", names[1:])
}
