package main

import (
	"fmt"
	"strconv"
	"strings"
)

func calculator(value1 string, value2 string) float64 {
	numValue1, err1 := strconv.ParseFloat(strings.TrimSpace(value1), 64)
	numValue2, err2 := strconv.ParseFloat(strings.TrimSpace(value2), 64)
	if err1 != nil || err2 != nil {
		panic("Invalid input. Please enter valid numbers.")
	} else {
		fmt.Println("The result of the calculation is: ", numValue1 + numValue2)
		return (numValue1 + numValue2)
	}
}