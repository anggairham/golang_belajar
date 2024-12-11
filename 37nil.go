package main

import "fmt"

func main() {
	var person map[string]string = nil
	fmt.Println(person)

	person2 := NewMap("Angga")
	fmt.Println(person2)
	person3 := NewMap("")
	if person3 == nil {
		fmt.Println("Data Kosong")
	}
}

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}

}
