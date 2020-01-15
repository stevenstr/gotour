/*
 *Author: Stefan
 *Date: 01/15/2020
 *Last changes: 01/15/2020 22:56
 *Task: Go tour
 */

package main

import (
	"fmt"
	"math"
)

//sqrt func
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

//min func
func main() {
	fmt.Println(sqrt(22), sqrt(-44))
}
