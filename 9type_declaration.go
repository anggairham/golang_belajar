package main

import "fmt"

// ini adalah alias
func main() {
	type NoKTP string
	type Married bool

	var noKTPAngga NoKTP = "123123123123"
	fmt.Println(noKTPAngga)
	var marriedStatus Married = true
	fmt.Println(marriedStatus)
}