package main

import "fmt"

// konstanta wajib ada value
// konstanta tidak ada complain/error jika tidak dipakai mungkin dipakai di file lain
// jika variabel ada error jika tidak dipakai
func main() {
	const firstName string = "Angga"
	const lastName = "Irham"
	const pi = 3.14

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(pi)

	// deklasri multiple constant
	const (
		country = "Indonesia"
		age     = 19
	)
	fmt.Println(country)
	fmt.Println(age)
}
