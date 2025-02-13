package linkedin

import "fmt"

func SlicesToObjects(colorNames []string, hexValues []int) []Color {
	// Your code goes here.
	// Create a slice of Color objects

	colors := make([]Color, 0)
	for k := range colorNames {
		color := Color{colorNames[k], hexValues[k]}
		colors = append(colors, color)
	}
	fmt.Println(colors)
	return colors
}

type Color struct {
	Name string
	Hex  int
}
