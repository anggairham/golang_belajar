package main

import "fmt"

// tipe data slice
// potongan dari data array
// ukuran slice bisa berubah
// slice adalah data yang mengakses sebagain/seluruh data di array
// 3 data: pointer,length,capacity
func main() {
	// array[low:hight]
	// array[low:]
	// array[:high]
	// array[:]

	var months = [...]string{
		"Januari",
		"Feb",
		"Mar",
		"Apr",
		"Mei",
		"Juni",
		"Juli",
		"Agus",
		"Sept",
		"Okt",
		"Nov",
		"Des",
	}

	var slice = months[4:7]
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	months[5] = "Diubah"
	fmt.Println(months)

	slice[0] = "Mei Lagi"
	fmt.Println(months)

	// append slice
	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	daySlice1 := days[5:]
	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Minggu Baru"
	fmt.Println(days)
	fmt.Println(daySlice1)

	daySlice2 := append(daySlice1, "Libur Baru")
	daySlice2[0] = "Ups"
	fmt.Println(daySlice2)
	fmt.Println(days)

	// make slice
	newSlice := make([]string, 2, 5)
	newSlice[0] = "Angga"
	newSlice[1] = "Angga"
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	// copy slice : memindahkan data
	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))
	copy(toSlice, fromSlice)
	fmt.Println(toSlice)

	// array vs slice
	iniArray := [...]int{1,2,3,4,5}
	iniSlice := []int{1,2,3,4,5}
	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
