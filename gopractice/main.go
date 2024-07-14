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
}