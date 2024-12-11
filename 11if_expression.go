package main

import "fmt"

func main() {
	name := "Angga"
	if name == "Angga" {
		fmt.Println("Hello Angga")
	} else if name == "Test" {
		fmt.Println("Hello Test")
	} else {
		fmt.Println("Wrong")
	}
	// if short statement
	// tamabhakn; length := len(name);
	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama dibolehkan")
	}
}
