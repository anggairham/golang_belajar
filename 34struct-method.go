package main

import "fmt"

func main() {
	rully := Customer2{Name: "rully"}
	rully.sayHi()
}

type Customer2 struct {
	Name, Address string
	Age           int
}

// struct method
func (customer Customer2) sayHi() {
	fmt.Println("Hello, My Name is", customer.Name)
}
