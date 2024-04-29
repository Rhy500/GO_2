package main

import "fmt"

func main() {
	var jam, menit int
	var member bool
	var total float64

	fmt.Println("Masukkan durasi sewa (jam dan menit):")
	fmt.Print("Jam: ")
	fmt.Scanln(&jam)
	fmt.Print("Menit: ")
	fmt.Scanln(&menit)
	fmt.Println("Apakah Anda member? (true/false):")
	fmt.Scanln(&member)

	hitungSewa(jam, menit, member, &total)

	fmt.Println(total)
}

func durasi(jam, menit int) int {
	if menit >= 10 {
		return jam + 1
	}
	return jam
}

func potongan(durasi int, tarif int) int {
	if durasi > 3 {
		return (durasi * tarif) / 10
	}
	return 0
}

func tarif(member bool) int {
	if member {
		return 3500
	}
	return 5000
}

func hitungSewa(jam, menit int, member bool, biaya *float64) {
	durasiSewa := durasi(jam, menit)
	tarifPerJam := tarif(member)
	totalSebelumDiskon := float64(durasiSewa * tarifPerJam)
	potonganDiskon := float64(potongan(durasiSewa, tarifPerJam))

	*biaya = totalSebelumDiskon - potonganDiskon
}
