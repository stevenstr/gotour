/*
 *Author: Stefan
 *Date: 01/26/2020
 *Last changes: 01/26/2020 01:25
 *Task: Go tour
 */

package main

import "fmt"

//main func
func main() {
	pow := make([]int, 10)

	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}

	fmt.Println()

	for index, _ := range pow {
		fmt.Printf("%d\n", index)
	}
}
