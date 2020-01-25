/*
 *Author: Stefan
 *Date: 01/25/2020
 *Last changes: 01/25/2020 16:38
 *Task: Go tour
 */

package main

import "fmt"

//main func
func main() {
	q := []int{2, 3, 6, 8, 13}

	r := []bool{true, true, false, true, true}

	ss := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, true},
		{6, false},
		{8, true},
		{13, true},
	}

	fmt.Println(q)
	fmt.Println(r)
	fmt.Println(ss)
	fmt.Println()

	fmt.Printf("q: %2v %6T \n\r\n", q, q)
	fmt.Printf("r: %2v %6T \n\r\n", r, r)
	fmt.Printf("ss: %2v %6T \n\r\n", ss, ss)
}
