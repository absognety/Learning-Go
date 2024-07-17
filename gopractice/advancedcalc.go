package main

import (
	"fmt"
	"strconv"
	"strings"
	"math"
)

func convertInputToValue(input string) float64 {
	result, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println("Error in parsing string to float: ", err)
		panic("Error in parsing string to float")
	}
	return result
}

func addValues(value1, value2 float64) float64 {
	return value1 + value2
}

func subtractValues(value1, value2 float64) float64 {
	return math.Abs(value1 - value2)
}

func multiplyValues(value1, value2 float64) float64 {
	return value1 * value2
}

func divideValues(value1, value2 float64) float64 {
	if value2 == 0 {
		fmt.Println("Cannot divide by zero")
		panic("Cannot divide by zero")
	}
	return value1 / value2
}

func calculate(input1 string, input2 string, operation string) float64 {
	// Convert input strings to float values
	value1 := convertInputToValue(input1)
	value2 := convertInputToValue(input2)

	// Perform the specified operation
	switch operation {
	case "+":
		return addValues(value1, value2)
	case "-":
		return subtractValues(value1, value2)
	case "*":
		return multiplyValues(value1, value2)
	case "/":
		return divideValues(value1, value2)
	default:
		fmt.Println("Invalid operation")
		panic("Invalid operation")
	}
}