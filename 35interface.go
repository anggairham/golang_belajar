package main

import "fmt"

/**
Interface = tipe data Abstract, tidak memilki implementasi langsung
interface berisikan definisi-definisi method
biasanya digunakan sebagai kontrak

# implementasi interface
- setiap tipe data yang sesuai dengan kontrak interface, secara otomatis dianggap sebagai interface tersebut
- sehingga tidak perlu mengimplementasikan interface secara manual
- hal ini berbeda dengan java dan php
*/
func main() {
	var angga Person
	angga.Name = "Angga"
	SayHello(angga)

	animal := Animal{Name: "Black"}
	SayHello(animal)
}

type HasName interface {
	GetName() string
}

func SayHello(name HasName) {
	fmt.Println("Hello", name.GetName())
}

type Person struct {
	Name string
}

// apapun itu yang mempunyai method  GetName, maka dia berhak mengikuti kontrak (interface)
func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}
