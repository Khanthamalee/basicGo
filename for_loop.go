package main

import "fmt"

func main() {
	fmt.Println("For_loop!!!!!")
	fmt.Println("Let 's fizzbuzz!!!!!")

	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("Fizzbuzz!", i)
		} else if i%3 == 0 && i%5 != 0 {
			fmt.Println("Fizz!", i)
		} else if i%3 != 0 && i%5 == 0 {
			fmt.Println("Buzz!", i)
		}
	}
}
