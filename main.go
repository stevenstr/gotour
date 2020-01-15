/*
 *Author: Stefan
 *Date: 01/15/2020
 *Last changes: 01/15/2020 20:05
 *Task: Go tour
 */

package main

import "fmt"

//bitl func
func bitl(sum int) (a, b int) {
	a = sum * 4 / 9
	b = sum - a
	return a, b
}

//main func
func main() {
	fmt.Println(bitl(10))
	fmt.Println(bitl(12))
	fmt.Println(bitl(17))
}
