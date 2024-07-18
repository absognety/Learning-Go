package main

import (
	"encoding/json"
	"strings"
)

// getCartFromJson() returns a slice containing cartItem objects.
func getCartFromJson(jsonString string) []cartItem {
    var carts = make([]cartItem, 0, 20)
	content := string(jsonString)
	decoder := json.NewDecoder(strings.NewReader(content))
	_, err := decoder.Token()
	if err != nil {
		panic(err)
	}
	var cartitem cartItem
	for decoder.More() {
		err = decoder.Decode(&cartitem)
		if err != nil {
			panic(err)
		}
		carts = append(carts, cartitem)
	}
	return carts
}