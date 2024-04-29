package main

import "fmt"

func main() {
	var b1, b2, b3, b4, hasilJumlah int
	fmt.Scan(&b1, &b2, &b3, &b4)
	hasilJumlah = penjumlahan(b1, b2, b3, b4)
	fmt.Println(hasilJumlah)
}
func penjumlahan(b1, b2, b3, b4 int) int {
	return (b1 + b2 + b3 + b4)
}
