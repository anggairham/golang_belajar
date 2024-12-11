package main

import "fmt"

// satu nama variabel unique hanya bisa satu tipe data
// kata kunci var tidak wajib bisa diganti dengan :=
func main() {
	var name string

	name = "Angga Irham"
	fmt.Println(name)

	name = "Angga Irham S"
	fmt.Println(name)
	// name = true

	var age = 22
	fmt.Println(age)
	// := deklarasi var diawal
	// simple deklarasi var dengan :=
	country := "Indonesia"
	fmt.Println(country)

	// deklarasi multiple variabel
	var (
		firstName = "Angga"
		lastName  = "Irham"
	)
	fmt.Println(firstName, lastName)

}
