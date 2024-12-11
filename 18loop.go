package main

import "fmt"

// perulangan hanya ada for loop 11 Des 2024
// tidak ada do while, while
func main() {
	counter := 1
	for counter <= 10 {
		fmt.Println("perulangan ke ", counter)
		counter++
	}
	// for dengan statement
	for i := 1; i <= 10; i++ {
		fmt.Println("statement ke ", i)
	}
	slice := []string{"Angga", "Irham", "Stianto"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}
	// for range untuk slice/array/map
	for i, value := range slice {
		fmt.Println("Index", i, "=", value)
	}
	// jika tidak terpakai index gunakan underscore _
	for _, value := range slice {
		fmt.Println(value)
	}
	person := make(map[string]string)
	person["name"] = "Angga"
	person["title"] = "Programmer"
	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}
