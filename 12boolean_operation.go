package main

import "fmt"

// && || !
func main() {
	var nilaiAkhir = 90
	absensi := 81

	var lulusAkhir = nilaiAkhir > 80
	var lulusAbsen = absensi > 80

	var lulus = lulusAbsen && lulusAkhir
	fmt.Println(lulusAkhir)
	fmt.Println(lulusAbsen)
	fmt.Println(lulus)
}
