/*
 *Author: Stefan
 *Date: 01/15/2020
 *Last changes: 01/15/2020 20:05
 *Task: Go tour
 */

package main

import "fmt"

func swap(a, b string) (x, y string) {
	return b, a
}

func main() {
	var a = "hello"
	var b = "world"

	fmt.Println(swap(a, b))
}
