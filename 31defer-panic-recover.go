package main

import "fmt"

/*
# Defer Function = function yang bisa kita jadwalkan untuk dieksekusi setelah sebuah function selesai di eksekusi
Defer Function akan dieksekusi walaupun terjadi error di function yang dieksekusi

# Panic Function = function yang digunakan untuk menghentikan program
Panic Function biasa diapnggil ketika terjadi error pada saat program berjalan
program akan terhenti, namun defer function tetap dieksekusi

# Recover Function : untuk menangkap data panic, dengan recover proses panic akan terhenti, dan tidak jadi terhenti karena panic
*/
func main() {
	// runApp(false)
	// runApp(true) //ketika error
	// runAppRecover(true)
	runAppRecover(false)
	runApplication(10)
	// runApplication(0) // ketika error
}

func logging() {
	fmt.Println("Selesai memamnggil function")
}

func runApplication(value int) {
	defer logging()
	fmt.Println("Run App")
	result := 10 / value
	fmt.Println("result", result)
}

func endApp() {
	fmt.Println("End App")
}
func runApp(error bool) {
	defer endApp()
	if error {
		panic("ERROR")
	}
	fmt.Println("App berjalan")
}

// recover harus ada pada fungsi defer
func endAppRecover() {
	message := recover()
	if message != nil {
		fmt.Println("Terjadi Error", message)
	}
	fmt.Println("End App recover")
}

func runAppRecover(error bool) {
	defer endAppRecover()
	if error {
		panic("runAppRecover ERROR")
	}
}
