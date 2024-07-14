package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	// ReadString will block until the delimiter is entered
	input, _ := reader.ReadString('\n')
	fmt.Println("You entered: ", input)
}