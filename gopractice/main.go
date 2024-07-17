package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	// ReadString will block until the delimiter is entered
	input, _ := reader.ReadString('\n')
	fmt.Println("You entered: ", input)


	fmt.Print("Enter a number: ")
	input, _ = reader.ReadString('\n')
	aFloat, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Value entered is: ", aFloat)
	}

	fmt.Println(calculator("23", "45"))
	fmt.Println(calculator("23.33", "45.89"))
	fmt.Println(calculator("10.0", "5.5"))

	fmt.Println(convertToMap([]string{"apple", "orange"}))
	fmt.Println(convertToMap([]string{"apple", "orange", "banana"}))

	cart := []cartItem{}
	var apples = cartItem{"apple", 2.99, 3}
	var oranges = cartItem{"orange", 1.50, 8}
	var bananas = cartItem{"banana", .49, 12}
	cart = append(cart, apples, oranges, bananas)
	fmt.Println(calculateTotal(cart))

	value1 := "10"
	value2 := "5.5"
	operation := "+"
	fmt.Println(calculate(value1, value2, operation))

}