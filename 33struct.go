package main

import "fmt"

/**
struct = sebuah template data yang digunakan untuk menggabungkan nol atau lebih tipe data lainnya dalam satu kesatuan
biasanya representrasi data dalam aplikasi yang dibuat
data di struct disimpan dalam field
struct adalah kumpulan dari field

# Struct Literals
*/
func main() {
	// membuat data dari struct
	var biasa Customer
	biasa.Name = "Angga"
	biasa.Address = "Indo"
	biasa.Age = 23
	fmt.Println(biasa)

	// # Struct Literals
	joko := Customer{
		Name: "Joko", Address: "Indo", Age: 22,
	}
	fmt.Println(joko)
	// ini harus urut sesuai posisi field
	budi := Customer{"Budi", "Jak", 30}
	fmt.Println(budi)
}

type Customer struct {
	Name, Address string
	Age           int
}
