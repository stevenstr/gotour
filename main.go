/*
 *Author: Stefan
 *Date: 01/24/2020
 *Last changes: 01/24/2020 19:41
 *Task: Go tour
 */

package main

import (
	"fmt"
	"time"
)

//main func
func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Morning!")
	case t.Hour() >= 12 || t.Hour() <= 19:
		fmt.Println("Day")
	case t.Hour() > 19:
		fmt.Println("Evening")
	default:
		fmt.Println("I don't know")
	}
}
