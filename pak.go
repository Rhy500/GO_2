package main

import (
	"fmt"
	"time"
)

// Variabel global
var (
	jenisKendaraan       string
	jam1, menit1, detik1 int
	jam2, menit2, detik2 int
	totalUang            int
)

// Menampilkan menu
func menu() {
	fmt.Println("--------------------------")
	fmt.Println("          M E N U         ")
	fmt.Println("--------------------------")
	fmt.Println("1. Input Kendaraan Masuk  ")
	fmt.Println("2. Input Kendaraan Keluar ")
	fmt.Println("3. Hitung Biaya Parkir    ")
	fmt.Println("4. Cetak Total Uang       ")
	fmt.Println("5. Exit                    ")
	fmt.Println("--------------------------")
}

// Prosedur input kendaraan masuk
func inputKendaraanMasuk() {
	fmt.Println("Masukkan jenis kendaraan (mobil/motor): ")
	fmt.Scan(&jenisKendaraan)
	fmt.Println("Masukkan jam, menit, detik kendaraan masuk (format: jam menit detik): ")
	fmt.Scan(&jam1, &menit1, &detik1)
}

// Prosedur input kendaraan keluar
func inputKendaraanKeluar() {
	fmt.Println("Masukkan jam, menit, detik kendaraan keluar (format: jam menit detik): ")
	fmt.Scan(&jam2, &menit2, &detik2)
}

// Prosedur hitung biaya parkir
func hitungBiayaParkir() {
	// Menghitung durasi parkir dalam jam
	waktuMasuk := time.Date(2024, time.Month(1), 1, jam1, menit1, detik1, 0, time.UTC)
	waktuKeluar := time.Date(2024, time.Month(1), 1, jam2, menit2, detik2, 0, time.UTC)
	durasiParkir := waktuKeluar.Sub(waktuMasuk).Hours()
	jamParkir := int(durasiParkir)

	// Menghitung biaya parkir berdasarkan jenis kendaraan
	var biayaParkir int
	switch jenisKendaraan {
	case "mobil":
		biayaParkir = 5000 + (jamParkir-1)*3000
	case "motor":
		biayaParkir = 2000 + (jamParkir-1)*1000
	default:
		fmt.Println("Jenis kendaraan tidak valid.")
		return
	}

	// Menambahkan biaya parkir ke total uang
	totalUang += biayaParkir

	// Mereset data kendaraan
	reset()
	fmt.Printf("Biaya parkir %s selama %d jam: Rp %d\n", jenisKendaraan, jamParkir, biayaParkir)
}

// Prosedur reset
func reset() {
	jenisKendaraan = ""
	jam1, menit1, detik1 = 0, 0, 0
	jam2, menit2, detik2 = 0, 0, 0
}

// Prosedur cetak total uang
func cetakTotalUang() {
	fmt.Printf("Total uang: Rp %d\n", totalUang)
}

func main() {
	var pilih int

	for {
		menu()
		fmt.Println("Pilih (1/2/3/4/5)? ")
		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			inputKendaraanMasuk()
		case 2:
			inputKendaraanKeluar()
		case 3:
			hitungBiayaParkir()
		case 4:
			cetakTotalUang()
		case 5:
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
