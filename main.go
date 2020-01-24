/*
 *Author: Stefan
 *Date: 01/24/2020
 *Last changes: 01/24/2020 21:08
 *Task: Go tour
 */

package main

import "fmt"

//card struct
type card struct {
	Name string
	Gen  string
}

//main func
func main() {
	vac := card{}

	vac.Name = "Goose"
	vac.Gen = "Man"

	fmt.Println(vac)
	fmt.Println(vac.Name)
	fmt.Println(vac.Gen)
}
