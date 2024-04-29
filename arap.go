package main

import (
	"fmt"
)

func main() {
	var qirat int
	fmt.Print("Masukkan jumlah uang dalam qirat: ")
	fmt.Scan(&qirat)

	dinar := qirat / 600
	qirat %= 600
	dirham := qirat / 60
	qirat %= 60
	fals := qirat / 6
	qirat %= 6

	fmt.Printf("\nHasil konversi:\n")
	fmt.Printf("Dinar: %d\n", dinar)
	fmt.Printf("Dirham: %d\n", dirham)
	fmt.Printf("Fals: %d\n", fals)
	fmt.Printf("Qirat: %d\n", qirat)
}
