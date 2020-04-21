package main

import "fmt"

func minDistance() {
	for {
		//enemy 1: name of enemy 1
		var enemy1 string
		fmt.Scan(&enemy1)

		//enemy 2: name of enemy 2
		var enemy2 string
		fmt.Scan(&enemy2)

		//distance from enemy 1
		var dist1 int
		fmt.Scan(&dist1)

		//distance from enemy 2
		var dist2 int
		fmt.Scan(&dist2)

		if dist1 < dist2 {
			fmt.Println(enemy1)
		} else {
			fmt.Println(enemy2)
		}
	}
}
