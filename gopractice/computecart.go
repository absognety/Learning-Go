package main

type cartItem struct {
	Name     string
	Price    float64
	Quantity int
}

func calculateTotal(cart []cartItem) float64 {
	var total float64 = 0.0
	for _, item := range cart {
		total += item.Price * float64(item.Quantity)
	}
	return total
}