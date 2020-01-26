/*
 *Author: Stefan
 *Date: 01/26/2020
 *Last changes: 01/26/2020 22:49
 *Task: Go tour
 */

package main

import "fmt"

//main func
func main() {
	m := make(map[int]string)

	for i := 1; i <= 5; i++ {
		m[i] = "txt"
	}

	a, ok := m[1]
	if ok {
		fmt.Println(a)
	} else {
		fmt.Println("not exist")
	}

	a, ok = m[2]
	if ok {
		fmt.Println(a)
	} else {
		fmt.Println("not exist")
	}

	//
	m[3] = "hhahaha"

	a, ok = m[3]
	if ok {
		fmt.Println(a)
	} else {
		fmt.Println("not exist")
	}

	a, ok = m[4]
	if ok {
		fmt.Println(a)
	} else {
		fmt.Println("not exist")
	}
	//
	delete(m, 4)

	a, ok = m[4]
	if ok {
		fmt.Println(a)
	} else {
		fmt.Println("not exist")
	}

}
