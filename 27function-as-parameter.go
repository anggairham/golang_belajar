package main

import "fmt"

func main() {
	sayHelloWithFilter("Angga", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("Kotor", filter)

	sayHelloWithFilter2("Irham", spamFilter)
}

func sayHelloWithFilter(name string, filter func(string) string) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func spamFilter(name string) string {
	if name == "Kotor" {
		return "..."
	} else {
		return name
	}
}

// function type declaration : solusi jika parameter terlalu panjang

type Filter func(string) string

func sayHelloWithFilter2(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}
