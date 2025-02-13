package linkedin

// Uncomment this import section to use required Go packages
import (
	"fmt"
	"strconv"
	"strings"
)


// calculate() returns the the result of adding 2 numbers
// after parsing them from strings
func Calculate(value1 string, value2 string) float64 {
	// Your code goes here.
	fmt.Println("Calculate simple")
	s1 := strings.TrimSpace(value1)
	s2 := strings.TrimSpace(value2)

	// Convert the first string to a float64
	v1, err := strconv.ParseFloat(s1, 64)
	if err != nil {
		fmt.Println("Error : ", err)
		panic("Value 1 must be a number")
	}

	// Convert the second string to a float64
	v2, err := strconv.ParseFloat(s2, 64)

	//handle the error returnned by parseFloat function.
	if err != nil {
		fmt.Println("Error : ", err)
		panic("Value 2 must be a number")
	}

	// Calculate and return the result
	result := v1 + v2

	return result
}
