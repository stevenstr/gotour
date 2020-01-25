/*
 *Author: Stefan
 *Date: 01/25/2020
 *Last changes: 01/25/2020 23:39
 *Task: Go tour
 */

package main

import "fmt"

//main func
func main() {
	sl := []int{}

	fmt.Printf("len: %5v | cap: %5v\n\r\n", len(sl), cap(sl))

	for i := 0; i <= 60; i++ {
		fmt.Printf("len: %5v | cap: %5v\n\r", len(sl), cap(sl))
		sl = append(sl, i)
	}
}
