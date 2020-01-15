/*
 *Author: Stefan
 *Date: 01/15/2020
 *Last changes: 01/15/2020 21:06
 *Task: Go tour
 */

package main

import "fmt"

//main func
func main() {

	var a int32 = 55
	var b int64 = 677

	var c int = int(a) + int(b)

	fmt.Println(a, b, " a+b=", c)

	fmt.Printf("a: Type %T  Value %v \n\n", a, a)
	fmt.Printf("b: Type %T  Value %v \n\n", b, b)
	fmt.Printf("c: Type %T  Value %v \n\n", c, c)
}
