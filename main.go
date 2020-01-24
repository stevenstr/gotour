/*
 *Author: Stefan
 *Date: 01/24/2020
 *Last changes: 01/24/2020 19:02
 *Task: Go tour
 */

package main

import (
	"fmt"
	"math"
)

//pow func
func pow(x, n, limit float64) float64 {
	if v := math.Pow(x, n); v < limit {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, limit)
		return limit
	}
}

//main func
func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 30),
	)
}
