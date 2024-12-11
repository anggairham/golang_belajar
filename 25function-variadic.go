package main

import "fmt"

func main() {
	total := sumAll(10, 10, 10, 10, 10, 1123)
	fmt.Println(total)

	// merubah slice ke varargs
	numbers := []int{10, 10, 213, 21, 111, 11}
	// cukup tambahkan ...
	total = sumAll(numbers...)
	fmt.Println(total)
}

/*
#Variadic Function
Variadic function adalah fungsi yang menggunakan fitur varargs
parameter yang berada di posisi terkahir, memiliki kemampuan dijadikan sebuah varargs
varargs(variable arguments) = datanya bisa menerima lebih dari satu input, atau anggap saja semacam array
apa bedanya dengan parameter biasa dengan tipe data Array?
	Jika parameter tipe array, kita wajib membuat array terlebih dahulu sebelum mengirimkan ke function
	Jika parameter menggunakan varargs, kita bisa langsung mengirim datanya, jika lebih dari satum cukup gunakan tanda koma
Ciri :
terdapat ...int pada parameter
variable arguments hanya boleh 1 di parameter, dan harus berada di paling akhir
*/

func sumAll(numbers ...int) int {
	// varargs numbers nanti berubah jadi slice disini
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

// ini error
// func test(number ...int, test string) {
// }
