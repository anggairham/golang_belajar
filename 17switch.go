package main

import "fmt"

func main() {
	name := "Angga"
	switch name {
	case "Angga":
		fmt.Println("Hello Angga")
	case "Irham":
		fmt.Println("Hello Irham")
	default:
		fmt.Println("default")
	}

	// switch short statement
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("true")
	case false:
		fmt.Println("false")
	}
	// switch tanpa kondisi
	length := len(name)
	switch {
	case length > 10:
		fmt.Println(">10")
	case length > 5:
		fmt.Println(">5")
	default:
		fmt.Println("default")

	}

}
