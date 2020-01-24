/*
 *Author: Stefan
 *Date: 01/24/2020
 *Last changes: 01/24/2020 21:51
 *Task: Go tour
 */

package main

import "fmt"

func main() {

	//create array
	var arr = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)

	//create slice
	var sl []int = arr[1:4]
	fmt.Println(sl)

	//automatic array len
	var arr1 = [...]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(arr1)
	fmt.Println(len(arr1))

}
