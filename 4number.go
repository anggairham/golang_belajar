package main

import "fmt"

// integer : int8,int16,int32,int64
// unsigned integer : uint8,uint16,uint32,uint64
// floating : float32,float64; jarang dipakai : complex64,complex128
// alias : bype = uint8, rune = int32, int = minimal int32 (tergantung OS, jika 32bit int32 jika 64bit int64), uint=minimal uint32

func main() {
	fmt.Println("Satu", 1)
	fmt.Println("Dua", 2)
	fmt.Println("Tiga Koma Lima", 3.5)
}
