package main

import (
	"fmt"
	"math/rand"
	"time"
)

func isLaunchSuccess() {
	var count = 10
	for count > 0 {
		fmt.Println(count)
		rand.Seed(time.Now().UnixNano())
		time.Sleep(time.Second)
		if rand.Intn(100) == 0 {
			break
		}
		count--
	}
	if count == 0 {
		fmt.Println("Liftoff!")
	} else {
		fmt.Println("Launch failed.")
	}
}
