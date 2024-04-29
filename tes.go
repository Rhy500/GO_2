package main

import (
	"fmt"
)

const nMax = 100

type barang struct {
	nama  string
	harga int
}

type keranjang [nMax]barang

func inputBelanja(arrKeranjang *keranjang, n *int) {
	fmt.Println("Masukkan jumlah barang:")
	fmt.Scan(n)

	for i := 0; i < *n; i++ {
		fmt.Println("Masukkan nama barang ke-", i+1, ":")
		var nama string
		fmt.Scan(&nama)

		fmt.Println("Masukkan harga barang ke-", i+1, ":")
		var harga int
		fmt.Scan(&harga)

		arrKeranjang[i] = barang{nama, harga}
	}
}

func tampilData(arrKeranjang *keranjang, n *int) {
	fmt.Println("\nDaftar barang yang dibeli:")
	totalHarga := 0
	for i := 0; i < *n; i++ {
		fmt.Printf("%d. %s (%d)\n", i+1, arrKeranjang[i].nama) //, arrKeranjang[i].harga)
		totalHarga += arrKeranjang[i].harga
	}

	fmt.Println("\nTotal belanja: ", totalHarga)
}

func main() {
	var keranjangan keranjang
	var jumlahBarang int

	inputBelanja(&keranjangan, &jumlahBarang)
	tampilData(&keranjangan, &jumlahBarang)
}
