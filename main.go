/*
 *Author: Stefan
 *Date: 01/15/2020
 *Last changes: 01/15/2020 22:44
 *Task: Go tour
 */

package main

import "fmt"

//main func
func main() {
	sum := 1

	for sum < 1000 {
		fmt.Println(sum)
		sum += sum
	}

	fmt.Println()
	fmt.Println()

	fmt.Println(sum)
}
