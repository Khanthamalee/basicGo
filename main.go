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

}
