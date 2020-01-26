/*
 *Author: Stefan
 *Date: 01/26/2020
 *Last changes: 01/26/2020 23:17
 *Task: Go tour
 */

package main

import "fmt"

//Bus type
type Bus struct {
	Storey int
	Place  int
	Res    int
}

//without pointers can't change source value
//Places method
func (b Bus) Places() int {
	pl := b.Storey * b.Place
	return int(pl)
}

//with pointer - can change source values

//PlacesPointer method
func (b *Bus) PlacesPointer() {
	b.Res = b.Storey * b.Place
}

//main func
func main() {
	bigBus := Bus{Storey: 2, Place: 33}
	fmt.Println(bigBus)

	b := bigBus.Places()
	fmt.Println(b)

	bigBus.PlacesPointer()
	fmt.Println(bigBus)
}
