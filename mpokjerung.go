package main

import "fmt"

// Konstanta untuk jumlah maksimum segitiga
const NMax int = 1000

// Definisi tipe array untuk menyimpan sisi segitiga
type tabSegitiga [NMax]int

func inputS(Tab1, Tab2, Tab3 *tabSegitiga, n int) {
	/*I.S. Terdefinisi bilangan bulat n.
	Proses: menerima masukan data sisi segitiga. Masukan berhenti apabila sudah
	mencapai nilai maksimal yaitu n.
	F.S. Array Tab1, Tab2, Tab3 berisi bilangan bulat yang diinputkan*/
	var i int
	for i = 0; i < n; i++ {
		fmt.Printf("Masukkan tiga sisi untuk segitiga ke-%d: ", i+1)
		fmt.Scan(&(*Tab1)[i], &(*Tab2)[i], &(*Tab3)[i])
	}
}

func CekSegitiga(Tab1, Tab2, Tab3 tabSegitiga, n int) {
	/*I.S. Terdefinisi bilangan bulat n. Kemudian Array Tab1, Tab2, dan Tab3
	yang berisi sisi segitiga sebanyak n.
	F.S. Menampilkan informasi jenis segitiga*/

	// Fungsi untuk menentukan jenis segitiga
	var i int
	for i = 0; i < n; i++ {

		// Tentukan jenis segitiga
		if Tab1[i] == Tab2[i] && Tab2[i] == Tab3[i] {
			fmt.Println("Segitiga ke-%d adalah segitiga sama sisi.\n", i+1)
		} else if Tab1[i] == Tab2[i] || Tab2[i] == Tab3[i] || Tab1[i] == Tab3[i] {
			fmt.Println("Segitiga ke-%d adalah segitiga sama kaki.\n", i+1)
		} else {
			fmt.Println("Segitiga ke-%d adalah segitiga sembarang.\n", i+1)
		}
	}
}
func main() {
	var TabS1, TabS2, TabS3 tabSegitiga
	var n int
	fmt.Scan(&n)
	inputS(&TabS1, &TabS2, &TabS3, n)
	CekSegitiga(TabS1, TabS2, TabS3, n)
}
