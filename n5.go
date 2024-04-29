package main

import "fmt"

const nMax int = 1000
type tabInt [nMax]int

func isiArray(arrInt *tabInt, n *int) {
    /*I.S. Data tersedia dalam piranti masukan
    F.S. arrInt telah terisi sejumlah bilangan 
    Proses akan berhenti ketika dimasukkan bilangan selain 1 dan 0*/
    var bil int
    fmt.Scan(&bil)
    for bil == 0 || bil == 1 && *n <= nMax{
        arrInt[*n] = bil
        *n = *n + 1
        fmt.Scan(&bil)
    }
}

func BinerToDesimal(arrInt tabInt, n int) int {
    /*mengembalikan nilai desimal yang telah dikonversi 
    dari representasi biner dalam array*/
    var pangkat,hasil int
    hasil = 0
    pangkat = 1
    for i := n - 1; i >= 0; i-- {
        hasil = hasil + arrInt[i]* pangkat
        pangkat = pangkat*2
    }
    return hasil
}

func main() {
	var A tabInt
	var N, desimal int
	isiArray(&A, &N)
	desimal = BinerToDesimal(A, N)
	fmt.Println(desimal)
}
