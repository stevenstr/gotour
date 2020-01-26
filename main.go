/*
 *Author: Stefan
 *Date: 01/26/2020
 *Last changes: 01/26/2020 22:34
 *Task: Go tour
 */

package main

import "fmt"

//Vertex type
type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

//main func
func main() {
	//just example
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	//create map
	h := make(map[int]string)

	//just joke
	for i := 0; i <= 10; i++ {
		if (i % 2) == 0 {
			h[i] = "NANANA"
		} else if (i % 2) != 0 {
			h[i] = "BARBAR"
		}
	}
	fmt.Println(h)

	//change value
	h[2] = "XAXAXA"
	fmt.Println(h)

	//get value from map
	a := h[1]
	b := h[3]
	c := h[5]

	fmt.Println(a, b, c)

	//check that some values contain or not
	_, ok := h[1]
	if !ok {
		fmt.Println("Map don't contain 1")
	}
	_, ok = h[4]
	if !ok {
		fmt.Println("Map don't contain 4")
	}
	_, ok = h[12]
	if !ok {
		fmt.Println("Map don't contain 12")
	}
}
