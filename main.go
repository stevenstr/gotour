/*
 *Author: Stefan
 *Date: 01/26/2020
 *Last changes: 01/26/2020 01:06
 *Task: Go tour
 */

package main

import "fmt"

//main func
func main() {
	arr := [6]int{0, 1, 4, 9, 16, 25}

	for i, v := range arr {
		fmt.Printf("%v * %v = %v\n\r", i, i, v)
	}
}
