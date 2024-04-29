package main

import "fmt"

func main() {
	var jam, menit, detik int
	fmt.Scan(&jam, &menit, &detik)
	fmt.Println("Hasil Konversi =", (konversiWaktu(jam, menit, detik)))
}
func konversiWaktu(jam, menit, detik int) int {
	return ((jam * 3600) + (menit * 60) + detik)
}
