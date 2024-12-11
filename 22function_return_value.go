package main

import "fmt"

func main() {
	sayHello := getHello("angga")
	fmt.Println(sayHello)
	fmt.Println(getHello(""))
}

// return value wajib menuliskan tipe data
func getHello(name string) string {
	if name == "" {
		return "Hello bro"
	} else {
		return "Hello " + name
	}
}
