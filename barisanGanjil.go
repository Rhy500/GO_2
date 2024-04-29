package main

import "fmt"

// Fungsi untuk menampilkan barisan bilangan ganjil dari 1 hingga N
func bilGanjil(N int) {
	if N < 1 {
		return
	} else {
		bilGanjil(N - 1)
	}

	if N%2 != 0 {
		fmt.Print(N, " ")
	}
}

func main() {
	var N int
	fmt.Scan(&N)

	fmt.Println(N)
	bilGanjil(N)

}
