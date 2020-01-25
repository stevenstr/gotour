/*
 *Author: Stefan
 *Date: 01/25/2020
 *Last changes: 01/25/2020 23:45
 *Task: Go tour
 */

package main

import "fmt"

func main() {
	var sl []int //are nil
	//var sl = []int{} // --- not are nil

	fmt.Printf("sl value: %4v | len: %4v | cap: %4v \n\r\n", sl, len(sl), cap(sl))

	if sl == nil {
		fmt.Println("Slice are nil!")
	}
}
