/*
 *Author: Stefan
 *Date: 01/24/2020
 *Last changes: 01/15/2020 19:32
 *Task: Go tour
 */

package main

import (
	"fmt"
	"time"
)

//main func
func main() {

	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}
