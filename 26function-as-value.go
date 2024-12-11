package main

import "fmt"

/*
function adalah first class citizen
function juga merupakan tipe data, dan bisa disimpan didalam variable
*/
func main() {
	goodbye := getGoodBye
	fmt.Println(goodbye("Angga"))
}

func getGoodBye(name string) string {
	return "Good Bye " + name
}
