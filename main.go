/*
 *Author: Stefan
 *Date: 01/24/2020
 *Last changes: 01/24/2020 23:05
 *Task: Go tour
 */

package main

import "fmt"

//main func
func main() {
	defer fmt.Println("!")
	defer fmt.Println("world")
	fmt.Println("Hello")
}
