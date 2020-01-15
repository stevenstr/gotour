/*
 *Author: Stefan
 *Date: 01/15/2020
 *Last changes: 01/15/2020 20:05
 *Task: Go tour
 */

package main

import "fmt"

var a int
var b float64
var C, text string

//main func
func main() {
	a = 5
	C = "C"
	b = 5.666
	text = "texttexttext"

	f := 55
	var g int = 5

	var n int
	var m string
	var x float64
	var k bool

	fmt.Printf("a:  Type %T  Value %v \n", a, a)
	fmt.Printf("C:  Type %T  Value %v \n", C, C)
	fmt.Printf("b:  Type %T  Value %v \n", b, b)
	fmt.Printf("text:  Type %T  Value %v \n\n", text, text)

	fmt.Printf("f:  Type %T  Value %v\n", f, f)
	fmt.Printf("g:  Type %T  Value %v\n\n", g, g)

	fmt.Printf("n:  Type %T  Value %v\n", n, n)
	fmt.Printf("m:  Type %T  Value %v\n", m, m)
	fmt.Printf("x:  Type %T  Value %v\n", x, x)
	fmt.Printf("k:  Type %T  Value %v\n", k, k)
}
