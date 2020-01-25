/*
 *Author: Stefan
 *Date: 01/26/2020
 *Last changes: 01/26/2020 00:49
 *Task: Go tour
 */

package main

import (
	"fmt"
)

//summs func
func summs(arr [6]int) (s int) {

	l := len(arr)
	sum := 0
	for i := 0; i < l; i++ {
		sum += arr[i]
	}
	return sum
}

//main func
func main() {
	arr1 := [6]int{1, 1, 1, 1, 1, 1}

	result := summs(arr1)
	fmt.Println(result)

	arr2 := [6]int{0, 1, 0, 1, 0, 1}

	result = summs(arr2)
	fmt.Println(result)
}
