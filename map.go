package main

import "fmt"

func main() {
	fmt.Println("Map!!")
	//สร้าง Map
	var country map[string]string = map[string]string{}
	fmt.Println("country :", country)
	//เพิ่มข้อมูลใน Map
	country["TH"] = "thailand"
	country["EN"] = "England"
	fmt.Println("country after update data :", country)
	//เรียกข้อมูลใน country
	fmt.Println("EN :", country["EN"])
	fmt.Println("TH :", country["TH"])
	fmt.Println("JP :", country["JP"])

}
