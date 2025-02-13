// package main

// import (
// 	"fmt"
// )

// func main() {
// 	//No input and no output
// 	showtopic()
// 	showmessage("No input and no output")
// 	//Have input and no output
// 	showmessage("Have input and no output")
// 	//No input and have output
// 	showmessage("No input and have output")
// 	showmvalue(sumincome())
// 	//Have input and have output
// 	showmessage("Have input and have output")
// 	var score int
// 	showmessage("Please input the score")
// 	fmt.Scanf("%d", &score)
// 	showmessage(grade(score))
// 	//Return multiple value
// 	fmt.Println(scoreandgrade(score))

// }

// // No input and no output
// func showtopic() {
// 	fmt.Println("Function")
// }

// // Have input and no output
// func showmessage(message string) {
// 	fmt.Println(message)
// }

// func showmvalue(value int) {
// 	fmt.Println(value)
// }

// // No input and have output
// func sumincome() int {
// 	number1 := 60
// 	number2 := 30
// 	total := number1 + number2
// 	return total
// }

// // Have input and have output
// func grade(score int) string {
// 	var g string
// 	if score >= 50 {
// 		g = " S"
// 	} else {
// 		g = " F"
// 	}
// 	return g
// }

// // Return multiple value
// func scoreandgrade(score int) (int,string) {
// 	var point int = score
// 	var g string
// 	if score >= 50 {
// 		g = " S"
// 	} else {
// 		g = " F"
// 	}
// 	return point,g
// }
