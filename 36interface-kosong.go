package main

import "fmt"

/**
interface kosong agar dapat menerima tipe data apapun
*/
func main() {
	// var data int = Ups() //error
	var data interface{} = Ups() //error
	fmt.Println(data)
	kosong := Ups()
	fmt.Println(kosong)
}

func Ups() interface{} {
	return 1
	return true
	return "Ups"
}
