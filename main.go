/*
 *Author: Stefan
 *Date: 01/15/2020
 *Last changes: 01/15/2020 22:39
 *Task: Go tour
 */

package main

import "fmt"

//main func
func main() {
	i := 0
	//0 1 2 3 4 5 6 7 8 9
	for i < 10 {
		fmt.Println(i)
		i++
	}

	//10 9 8 7 6 5 4 3 2 1 0
	for i >= 0 {
		fmt.Println(i)
		i--
	}

	fmt.Println()
	fmt.Println()

	// 1 4 6 16 25 36 49 64 81 100
	for x := 1; x <= 10; x++ {
		fmt.Println(x * x)
	}

	fmt.Println()
	fmt.Println()

	// 10 9 8 Continue 6 5 Continue 3 Continue 1 Break
	a := 10

	for {
		if a == 0 {
			fmt.Println("Break")
			break
		}
		if a == 7 {
			fmt.Println("Continue")
			a--
			continue
		}
		if a == 4 {
			fmt.Println("Continue")
			a--
			continue
		}
		if a == 2 {
			fmt.Println("Continue")
			a--
			continue
		}
		fmt.Println(a)
		a--
	}
}
