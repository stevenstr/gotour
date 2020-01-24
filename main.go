/*
 *Author: Stefan
 *Date: 01/24/2020
 *Last changes: 01/24/2020 21:44
 *Task: Go tour
 */

package main

import "fmt"

//main func
func main() {
	//array cannot be resized
	var a [2]string

	a[0] = "Hello "
	a[1] = "world!"

	fmt.Printf("a[0]: %7v %7T \n", a[0], a[0])
	fmt.Printf("a[1]: %7v %7T \n\n\r", a[1], a[1])

	fmt.Println(a)

	var c = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := 0; i < len(c); i++ {
		fmt.Printf("|i: %5v |c[i]: %5v|\n", i, c[i])
	}
}
