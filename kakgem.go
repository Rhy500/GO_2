package main

import "fmt"

const nMax int = 100

type belanja struct {
	nama  string
	harga int
}
type keranjang [nMax]belanja

func inputBelanja(arrKeranjang *keranjang, n *int) {
	/*I.S. Terdefinisi array arrKeranjang dan bilangan bulat n*/
	/*F.S. Array arrKeranjang terisi dengan data yang telah diinputkan*/
	fmt.Print("Masukkan jumlah barang yang akan dibeli: ")
	fmt.Scan(n) // Membaca jumlah barang

	for i := 0; i < *n; i++ {
		var nama string
		var harga int

		// Menginput nama barang dan harga
		fmt.Printf("Masukkan nama dan harga barang ke-%d: ", i+1)
		fmt.Scan(&nama, &harga)

		// Mengisi keranjang belanja
		arrKeranjang[i].nama = nama
		arrKeranjang[i].harga = harga
	}
}

func tampilData(arrKeranjang *keranjang, n *int) {
	/*I.S. Terdefinisi array arrKeranjang dan bilangan bulat n*/
	/*F.S. Menampilkan daftar barang yang dibeli beserta total harga dari barang-barang tersebut*/
	var totalHarga int

	fmt.Println("Daftar barang yang dibeli:")
	for i := 0; i < *n; i++ {
		fmt.Printf("%d. %s\n", i+1, arrKeranjang[i].nama)
		totalHarga += arrKeranjang[i].harga // Menghitung total harga
	}

	// Menampilkan total harga belanja
	fmt.Printf("Total belanja: %d\n", totalHarga)
}
func main() {
	var arrKeranjang keranjang
	var n int
	inputBelanja(&arrKeranjang, &n)
	tampilData(&arrKeranjang, &n)
}
