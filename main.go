package main

import (
	"basicGo/linkedin"
	"fmt"
)

func main() {
	//time
	linkedin.Time()
	fmt.Println("-------------------------")
	v1, v2 := "10", "5.5"
	fmt.Println(linkedin.Calculate(v1, v2))
	fmt.Println("-------------------------")
	linkedin.Pointer()
	fmt.Println("-------------------------")
	colorNames := []string{"Red", "Green", "Blue"}
	hexValues := []int{0xFF0000, 0x00FF00, 0x0000FF}
	linkedin.SlicesToObjects(colorNames, hexValues)
	fmt.Println("-------------------------")
	var cart []linkedin.CartItem
	var apples = linkedin.CartItem{Name: "apple", Price: 2.99, Quantity: 3}
	var oranges = linkedin.CartItem{Name: "orange", Price: 1.50, Quantity: 8}
	var bananas = linkedin.CartItem{Name: "banana", Price: .49, Quantity: 12}
	cart = append(cart, apples, oranges, bananas)
	fmt.Println(linkedin.CalculateTotal(cart))

}
