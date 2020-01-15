/*
 *Author: Stefan
 *Date: 01/15/2020
 *Last changes: 01/15/2020 21:40
 *Task: Go tour
 */

package main

import "fmt"

const (
	//Pi const
	Pi = 1
	//Di const
	Di
	//Ci const
	Ci = 2
	//Fi const
	Fi
)

const (
	//Host const
	Host = ":8000"
	//IP const
	IP = "127.0.0.1"
	//User const
	User = "postgres"
)

//main func
func main() {

	fmt.Println("\tSome constant:")
	fmt.Printf("Pi: Type %8T Value %8v\n\n", Pi, Pi)
	fmt.Printf("Di: Type %8T Value %8v\n\n", Di, Di)
	fmt.Printf("Ci: Type %8T Value %8v\n\n", Ci, Ci)
	fmt.Printf("Fi: Type %8T Value %8v\n\n", Fi, Fi)
	fmt.Printf("Host: Type %8T Value %8v\n\n", Host, Host)
	fmt.Printf("IP: Type %8T Value %8v\n\n", IP, IP)
	fmt.Printf("User: Type %8T Value %8v\n\n", User, User)
}
