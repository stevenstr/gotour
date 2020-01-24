/*
 *Author: Stefan
 *Date: 01/24/2020
 *Last changes: 01/24/2020 20:58
 *Task: Go tour
 */

package main

import "fmt"

//card struct
type card struct {
	Number string
	Date   string
}

//main func
func main() {

	vac := card{"3334455", "ammmm"}
	vac1 := card{"5436677", "hmmmm"}

	fmt.Println(vac, vac1)
	fmt.Println()
	fmt.Println(vac.Number, vac.Date)
	fmt.Println()
	fmt.Println(vac1.Number, vac1.Date)
}
