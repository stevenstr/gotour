/*
 *Author: Stefan
 *Date: 01/24/2020
 *Last changes: 01/24/2020 22:14
 *Task: Go tour
 */

package main

import "fmt"

//main func
func main() { //    0  1  2  3  4  5  6  7  8  9
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Source array: ", arr)

	//create two slices to show how working referencing
	sl1 := arr[2:5] // start with 2 and to 5 (2,3,4) five not included
	sl2 := arr[3:8]

	fmt.Println()
	fmt.Println("Slice 1: ", sl1) // 3, 4, 5
	fmt.Println("Slice 2: ", sl2) // 4, 5, 6, 7, 8

	//change the first slice
	fmt.Println(sl1[2])
	sl1[2] = 555 //change 5 on 555
	fmt.Println(sl1[2])

	fmt.Println("\nChanged values\nSource array: ", arr)
	fmt.Println("Slice 1: ", sl1)
	fmt.Println("Slice 2: ", sl2)
}
