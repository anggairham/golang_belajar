package main

import "fmt"

// > < >= <= == !=
func main()  {
	var name1 = "Angga"
	var name2 ="Angga"
	var result bool = name1 == name2
	fmt.Println(result)

	value1 := 100
	value2 := 200
	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)
}