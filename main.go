/*
 *Author: Stefan
 *Date: 01/24/2020
 *Last changes: 01/24/2020 20:17
 *Task: Go tour
 */

package main

import "fmt"

//main func
func main() {
	fmt.Println("Start from 10")

	for i := 0; i < 11; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("End to 0")
}
