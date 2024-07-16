package main

func convertToMap(items []string) map[string]float64 {
	numofitems := len(items)
	equalprice := 100.0 / float64(numofitems)
	result := make(map[string]float64)
	for i := 0; i < numofitems; i++ {
		result[items[i]] = equalprice
	}
	return result
}