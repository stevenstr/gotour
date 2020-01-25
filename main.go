/*
 *Author: Stefan
 *Date: 01/25/2020
 *Last changes: 01/25/2020 23:58
 *Task: Go tour
 */

package main

import "fmt"

//main func
func main() {
	sl := make([]int, 5)
	prSl(sl)
	sl1 := make([]int, 5, 5)
	prSl(sl1)
	sl2 := make([]int, 0, 5)
	prSl(sl2)
	sl3 := make([]int, 10, 10)
	prSl(sl3)
	sl4 := sl3[3:6]
	prSl(sl4)
}

//prSl func
func prSl(sl []int) {
	fmt.Printf("sl: %2v | len: %4v | cap: %4v \n\r\n", sl, len(sl), cap(sl))
}
