/*
 *Author: Stefan
 *Date: 01/15/2020
 *Last changes: 01/15/2020 20:05
 *Task: Go tour
 */

package main

import "fmt"

var a, b, c int = 12, 13, 14

//main func
func main() {
	var text, ip string = "text", "127.0.0.1"
	host := ":8000"
	user, host := "vasia", ":8080"

	fmt.Println(a, b, c)
	fmt.Println(text, ip, host, user)
}
