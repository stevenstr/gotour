/*
 *Author: Stefan
 *Date: 01/15/2020
 *Last changes: 01/15/2020 23:05
 *Task: Go tour
 */

package main

import (
	"fmt"
	"math"
)

//pow func
func pow(a, b, limit float64) float64 {
	if val := math.Pow(a, b); val < limit {
		return val
	}
	return limit
}

//main func
func main() {
	fmt.Println(pow(3, 4, 81), pow(3, 2, 10))
}
