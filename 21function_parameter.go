package main

import "fmt"

func main() {
	sayHelloTo("Angga", "Irham")
}

func sayHelloTo(firstName string, lastName string) {
	fmt.Println(firstName, lastName)
}
