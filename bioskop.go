package main

import "fmt"

type Tiket struct {
	film, hari string
	tglmer     bool
}

func main() {
	// Contoh penggunaan
	var film, hari string
	var tglmer bool
	//var tiket Tiket
	fmt.Scan(&film, &hari, &tglmer)
	//tglmer = true
	tiket := Tiket{film, hari, tglmer}
	biaya := hitungHarga(tiket)
	fmt.Println("Biaya yang harus dibayar:", biaya)
}
func hitungHarga(tiket Tiket) int {
	hargaTiket := 50000 // Harga tiket hari kerja
	if tiket.hari == "Sabtu" || tiket.hari == "Minggu" || tiket.tglmer {
		hargaTiket += hargaTiket * 50 / 100 // Tambah 50% jika akhir pekan atau tanggal merah
	} else {
		hargaTiket = hargaTiket
	}
	return hargaTiket
}
