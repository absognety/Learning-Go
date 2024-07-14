package main

import "fmt"

func torch() {
	var hasTorch = true
	var litTorch = false

	if !hasTorch || !litTorch {
		fmt.Println("Nothing to see here!")
	}
}
