package main
import "fmt"
const nMax int = 1000
type tabInt [nMax]int

func isiArray(arr *tabInt, n int) {
	/* I.S. Data tersedia dalam piranti masukan
	F.S. arr telah terisi sejumlah bilangan */
	var bil, i int
	i = 0
	for i < n {
		fmt.Scan(&bil)
		for bil < 1 || bil > 150 {
			fmt.Scan(&bil)
		}
		arr[i] = bil
		i = i + 1
	}
}

func pembagianPermen(arr tabInt, n int) tabInt {
	/*mengembalikan array yang berisi jumlah pembagian permen untuk 
	setiap murid*/
	var permen tabInt
	var i int

	if n == 1 {
		permen[0] = 1
	} else {
		for i = 0; i < n; i++ {
			permen[i] = 1
		}
		i = 0
		for i < n {
			if i > 0 && arr[i] > arr[i-1] {
				permen[i] = permen[i-1] + 1
			}
			i = i + 1
		}
		i = n - 1
		for i > 0 {
			if arr[i-1] > arr[i] && permen[i-1] <= permen[i] {
				permen[i-1] = permen[i] + 1
			}
			i = i - 1
		}
	}
	return permen
}
func permenMinimum(arr tabInt, n int) int {
	/* Mengembalikan permen minimum yang harus dimiliki Raisa agar semua
	murid mendapatkan permen semua dan sesuai dengan ketentuan */
	var total, i int
	total = 0
	for i = 0; i < n; i++ {
		total = total + arr[i]
	}
	return total
}

func cetak(arr tabInt, n int) {
	var i, total int
	var permen tabInt
	permen = pembagianPermen(arr, n)
	for i = 0; i < n; i++ {
		fmt.Print(permen[i]," ")
	}
	total = permenMinimum(permen, n)
	fmt.Println()
	fmt.Print(total)
}

func main() {
	var N int
	var rank tabInt
	fmt.Scan(&N)
	isiArray(&rank, N)
	cetak(rank, N)
}