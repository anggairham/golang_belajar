package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		if i == 6 {
			break
		}
		fmt.Println("perulangan ke", i)
	}
}
