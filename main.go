/*
 *Author: Stefan
 *Date: 01/24/2020
 *Last changes: 01/24/2020 20:17
 *Task: Go tour
 */

package main

import "fmt"

type card struct {
	Name string
	Data string
}

func main() {

	vac := card{}

	vac.Name = "Stas"
	vac.Data = "Nan"

	field := &vac.Name

	fmt.Printf("vac: %9v %8T\n\n", vac, vac)
	fmt.Printf("vac.Name: %9v %8T\n\n", vac, vac)
	fmt.Printf("vac.Data: %9v %8T\n\n", vac.Data, vac.Data)
	fmt.Printf("vac.Name: %9v %8T\n\n", vac.Name, vac.Name)

	fmt.Printf("field as pointer: %5v %5T  field as value: %5v %5T\n\n", field, field, *field, *field)
}
