/*
 *Author: Stefan
 *Date: 01/24/2020
 *Last changes: 01/24/2020 23:31
 *Task: Go tour
 */

package main

import "fmt"

//main func
func main() {
	sl := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("sl values: %v|type: %4T|len: %4v| cap: %4v\n\r\n", sl, sl, len(sl), cap(sl))
	sl = sl[1:5]
	fmt.Printf("sl values: %v|type: %4T|len: %4v| cap: %4v\n\r\n", sl, sl, len(sl), cap(sl))
	sl = sl[1:]
	fmt.Printf("sl values: %v|type: %4T|len: %4v| cap: %4v\n\r\n", sl, sl, len(sl), cap(sl))
	sl = append(sl, 55)
	fmt.Printf("sl values: %v|type: %4T|len: %4v| cap: %4v\n\r\n", sl, sl, len(sl), cap(sl))
}
