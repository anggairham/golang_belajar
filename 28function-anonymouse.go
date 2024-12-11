package main

import "fmt"

// function tanpa nama
func main() {
	blacklist := func(name string) bool {
		return name == "admin"
	}
	registerUser("angga", blacklist)
	registerUser("irham", func(name string) bool {
		return name == "irham"
	})
}

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

// func blacklistAdmin(name string) bool {
// 	return name == "admin"
// }
// func blacklistRoot(name string) bool {
// 	return name == "root"
// }
