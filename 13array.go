package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Angga"
	names[1] = "Irham"
	names[2] = "S"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	// membuat array langsung
	var values = [3]int{
		80,90,100,
	}
	fmt.Println(values)

	// array function
	fmt.Println(len(values))
	values[0] = 10
	fmt.Println(values[0])
}