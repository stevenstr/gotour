/*
 *Author: Stefan
 *Date: 01/15/2020
 *Last changes: 01/15/2020 21:53
 *Task: Go tour
 */

package main

import "fmt"

const (
	Big    = 1 << 100
	PreBig = 1 << 50

	Small    = Big >> 99
	PreSmall = Big >> 59
)

//needInt func
func needInt(x int) int {
	return x*10 + 1
}

//needFloat func
func needFloat(x float64) float64 {
	return x * 0.1
}

//main func
func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(PreSmall))
	fmt.Println(needFloat(Big))
	fmt.Println(needFloat(PreBig))
}
