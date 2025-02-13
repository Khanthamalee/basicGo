package linkedin

import (
	"encoding/json"
	"strings"
)

func GetCartFromJson(jsonString string) []CartItem {
	var cart []CartItem
	// Your code goes here.
	decoder := json.NewDecoder((strings.NewReader(jsonString)))
	_, err := decoder.Token()
	checkError(err)
	var cartItem CartItem
	for decoder.More() {
		err := decoder.Decode(&cartItem)
		checkError(err)
		cart = append(cart, cartItem)
	}
	return cart
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
