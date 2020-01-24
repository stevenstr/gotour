/*
 *Author: Stefan
 *Date: 01/15/2020
 *Last changes: 01/15/2020 23:05
 *Task: Go tour
 */

package main

import (
	"fmt"
	"runtime"
)

//main func
func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}
}
