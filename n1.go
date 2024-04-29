package main

import "fmt"

type biaya struct {
	kendaraan string
	durasi    int
}

func main() {
	var parkir biaya
	hitungTarif(&parkir)
}
func hitungTarif(parkir *biaya) {
	var pMotor int = 1000
	var pMobil int = 5000
	var total int = 0
	fmt.Scan(&parkir.kendaraan, &parkir.durasi)
	for (parkir.kendaraan == "motor" || parkir.kendaraan == "mobil") && (parkir.durasi >= 0) {
		var biaya int = 0
		if parkir.kendaraan == "motor" {
			biaya = pMotor * parkir.durasi
		} else if parkir.kendaraan == "mobil" {
			biaya = pMobil * parkir.durasi
		}
		total = total + biaya
		fmt.Scan(&parkir.kendaraan, &parkir.durasi)
	}
	fmt.Println(total)
}