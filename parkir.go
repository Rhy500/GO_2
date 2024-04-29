package main

import "fmt"

func main() {
	var jenis string
	var jam, menit int
	fmt.Scan(&jenis, &jam, &menit)
	fmt.Print(hitungongkos(jenis, jam, menit))
}
func hitungongkos(jenis string, jam, menit int) int {
	var ongkos int
	tarif_mobil_biasa_perjam_pertama := 5000
	tarif_mobil_biasa_perjam_selanjutnya := 6000
	tarif_truk_perjam_pertama := tarif_mobil_biasa_perjam_pertama * 2
	tarif_truk_perjam_selanjutnya := tarif_mobil_biasa_perjam_selanjutnya * 2
	total_jam := jam
	if menit >= 10 {
		total_jam = total_jam + 1
	}

	if jenis == "truk" {
		if total_jam <= 2 {
			ongkos = total_jam * tarif_truk_perjam_pertama
		} else {
			ongkos = (2 * tarif_truk_perjam_pertama) + ((total_jam - 2) * tarif_truk_perjam_selanjutnya)
		}
	} else {
		if total_jam <= 2 {
			ongkos = total_jam * tarif_mobil_biasa_perjam_pertama
		} else {
			ongkos = (2 * tarif_mobil_biasa_perjam_pertama) + ((total_jam - 2) * tarif_mobil_biasa_perjam_selanjutnya)
		}
	}
	return ongkos
}
