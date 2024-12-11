package main

import "fmt"

func main() {
	firstName, middleName, lastName := getFullName()
	fmt.Println(firstName, lastName)
	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)

	// ignore beberapa value
	test, _, _ := getFullName()
	fmt.Println(test)
}

func getFullName() (string, string, string) {
	return "Angga", "Irham", "Stianto"
}
