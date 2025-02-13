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

}
