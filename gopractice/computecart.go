package main

type cartItem struct {
	name     string
	price    float64
	quantity int
}

func calculateTotal(cart []cartItem) float64 {
	var total float64 = 0.0
	for _, item := range cart {
		total += item.price * float64(item.quantity)
	}
	return total
}