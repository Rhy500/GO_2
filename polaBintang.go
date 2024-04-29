package main

import (
	"fmt"
)

// Fungsi rekursif untuk menampilkan pola bintang
func bintang(n int) {
	// Base case: jika n kurang dari 1, berhenti
	if n < 1 {
		return
	} else {

		// Panggil rekursif untuk menampilkan pola bintang sebanyak n-1
		bintang(n - 1)
	}

	// Cetak pola bintang sebanyak n
	for i := 0; i < n; i++ {
		fmt.Print("*")
	}
	fmt.Println()
}

func main() {
	var N int
	fmt.Print("Masukkan bilangan bulat positif N: ")
	fmt.Scan(&N)

	fmt.Println("Pola bintang dari 1 hingga", N, "adalah:")
	bintang(N)
}
