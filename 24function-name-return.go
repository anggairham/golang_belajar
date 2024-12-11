package main

import "fmt"

func main() {
	a, b, c := getCompleteName()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

}

// ini jarang di beberapa bahasa lain
// jika semua return tipenya sama cukup sekali, jika beda beda definisikan satu satu
func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Angga"
	middleName = "Irham"
	lastName = "Stianto"

	// ada 2 cara, bisa di tuliskan atau langsung return
	// return firstName, middleName, lastName
	return
}
