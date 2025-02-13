package linkedin

// Uncomment this import section to use required Go packages

type CartItem struct {
	Name     string
	Price    float64
	Quantity int
}

func CalculateTotal(cart []CartItem) float64 {
	// Your code goes here.
	var sum float64
	for i := range cart {
		amount := cart[i].Price * float64(cart[i].Quantity)
		sum = sum + amount
		//fmt.Println(sum)
	}
	return sum
}
