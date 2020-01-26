/*
 *Author: Stefan
 *Date: 01/26/2020
 *Last changes: 01/26/2020 22:38
 *Task: Go tour
 */

package main

import "fmt"

//Vertex type
type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

//main func
func main() {
	fmt.Println(m)
}
