/*
 *Author: Stefan
 *Date: 01/24/2020
 *Last changes: 01/24/2020 20:42
 *Task: Go tour
 */

package main

import "fmt"

func main() {
	i, j := 40, 2407

	p := &i //get address

	fmt.Printf("i: %4v %4T   p as pointer: %4v %4T  p as value:  %4v %4T\n\n", i, i, p, p, *p, *p)

	*p = 25 //change i value through dereferencing pointer
	fmt.Printf("i: %4v %4T   p as pointer: %4v %4T  p as value:  %4v %4T\n\n", i, i, p, p, *p, *p)

	p = &j // get address
	fmt.Printf("j: %4v %4T   p as pointer: %4v %4T  p as value:  %4v %4T\n\n", j, j, p, p, *p, *p)
	fmt.Printf("i: %4v %4T\n\n", i, i)

	*p = 7 //change j value thruogh dereferencing pointer
	fmt.Printf("j: %4v %4T  p as pointer: %4v %4T  p as value: %4v %4T\n\n", j, j, p, p, *p, *p)
}
