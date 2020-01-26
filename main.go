/*
 *Author: Stefan
 *Date: 01/26/2020
 *Last changes: 01/26/2020 23:01
 *Task: Go tour
 */

package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.

//fibonacci func
func fibonacci(x int) {
	f0 := 0
	f1 := 1
	f2 := 1
	fx := 0
	if x == 0 {
		fmt.Println(f0)
		return
	}
	if x == 1 {
		fmt.Println(f0)
		fmt.Println(f1)
		return
	}
	if x == 2 {
		fmt.Println(f0)
		fmt.Println(f1)
		fmt.Println(f2)
		return
	}

	fmt.Println(f0)
	fmt.Println(f1)
	fmt.Println(f2)

	for i := 3; i <= x; i++ {
		fx = f1 + f2
		fmt.Println(fx)
		f1 = f2
		f2 = fx
	}
}

//main func
func main() {
	fibonacci(12)
}
