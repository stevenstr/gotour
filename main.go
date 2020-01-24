/*
 *Author: Stefan
 *Date: 01/24/2020
 *Last changes: 01/24/2020 20:29
 *Task: Go tour
 */

package main

import "fmt"

//card struct
type card struct {
	Name string
	Data string
}

//main func
func main() {
	vac := card{Name: "Goose", Data: "something"}
	vac1 := card{}
	vac2 := card{Name: "Goodz"}

	fmt.Println(vac)
	fmt.Println(vac1)
	fmt.Println(vac2)
}
