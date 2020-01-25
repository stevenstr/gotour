/*
 *Author: Stefan
 *Date: 01/26/2020
 *Last changes: 01/26/2020 00:55
 *Task: Go tour
 */

package main

import "fmt"

//main func
func main() {
	var sl []int
	prSl(sl)

	sl = append(sl, 1)
	prSl(sl)
	sl = append(sl, 1)
	prSl(sl)
	sl = append(sl, 1)
	prSl(sl)
	sl = append(sl, 1)
	prSl(sl)
	sl = append(sl, 1)
	prSl(sl)
	sl = append(sl, 1)
	prSl(sl)
}

//prSl func
func prSl(sl []int) {
	fmt.Printf("sl: %v | len: %3v | cap: %3v \n\r\n", sl, len(sl), cap(sl))
}
