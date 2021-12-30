/*
Refer this example from Get Programming with Go by Nathan youngman and Roger Peppe
*/
package main

import (
	"fmt"
)

func determineLeapYear() {
	var year = 1900
	if ((year%4 == 0) && (year%100 != 0)) || (year%400 == 0) {
		fmt.Printf("Given year %v is a leap year", year)
	} else {
		fmt.Printf("Given year %v is not a leap year", year)
	}
}
