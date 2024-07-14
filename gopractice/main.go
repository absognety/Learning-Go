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
}