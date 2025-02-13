package linkedin

import (
	"fmt"
	"strconv"
	"strings"
)

func CalculateAdvance(input1 string, input2 string, operation string) float64 {
	// Your code goes here.
	var result float64
	value1 := convertInputToValue(input1)
	value2 := convertInputToValue(input2)

	switch operation {
	case "+":
		result = addValues(value1, value2)
	case "-":
		result = subtractValues(value1, value2)
	case "*":
		result = multiplyValues(value1, value2)
	case "/":
		result = divideValues(value1, value2)
	default:
		panic("Invalid operation")
	}
	return result
}

func convertInputToValue(input string) float64 {
	s1 := strings.TrimSpace(input)
	v1, err := strconv.ParseFloat(s1, 64)
	if err != nil {
		message := fmt.Sprintf(" %v Error : %v ", input, err)
		panic(message)
	}
	return v1
}

func addValues(value1, value2 float64) float64 {
	return value1 + value2
}

func subtractValues(value1, value2 float64) float64 {
	return value1 - value2
}

func multiplyValues(value1, value2 float64) float64 {
	return value1 * value2
}

func divideValues(value1, value2 float64) float64 {
	return value1 / value2
}
