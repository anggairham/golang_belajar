package main

import (
	"errors"
	"fmt"
)

func main() {
	var contohError error = errors.New("Ups Error")
	fmt.Println(contohError.Error())

	hasil, err := Pembagian(100, 10)
	if err == nil {
		fmt.Println("Hasil", hasil)
	} else {
		fmt.Println("Error", err.Error())
	}
}

// karena type error adalah interface boleh nil
func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh 0")
	} else {
		result := nilai / pembagi
		return result, nil
	}

}
