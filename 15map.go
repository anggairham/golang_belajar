package main

import "fmt"

// map bisa digunakan untuk menentukan tipe data index data
// map = key - value
// jumlah map bebas
func main() {
	var person2 map[string]string = map[string]string{
		"name": "Angga",
	}
	fmt.Println(person2)
	// cara singkat
	// map[tipe_key]tipe_value
	person := map[string]string{
		"name":    "Angga",
		"country": "Indo",
	}
	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["country"])
	person["name"] = "Irham"
	fmt.Println(person["name"])
	person["title"] = "Programmer"
	fmt.Println(person)
	fmt.Println(len(person))
	// membuat map tanpa data apapun dengan make
	// var book map[string]string = make(map[string]string)
	book := make(map[string]string)
	book["title"] = "Belajar golang"
	book["author"] = "Angga"
	book["response"] = "Ok"
	fmt.Println(book)
	delete(book, "response")
	fmt.Println(book)
}
