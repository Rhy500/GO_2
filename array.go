package main 

import "fmt"

//Sebuah tipe array of integer yang berkapasitas 256!
const kapasitas = 256

type arrayInt [kapasitas]int

//Buatlah procedure untuk pengisian array tersebut dengan sejumlah n bilangan.
func isiArray(arr *arrayInt, n int) {
	fmt.Println("Masukkan", n, "bilangan:")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
}

//Buatlah procedure untuk reverse isi dari array!
func reverseArray(arr *arrayInt) {
	for i := 0; i < kapasitas/2; i++ {
		arr[i], arr[kapasitas-1-i] = arr[kapasitas-1-i], arr[i]
	}
}

//Buatlah function untuk mengecek apakah suatu array memiliki pola palindrom! Nilai elemen membentuk pola simetris. ​
Contoh A = [10, 20, 30, 20, 10], B = [15, 75, 75, 15] dan C = [100]
func isPalindrom(arr *arrayInt) bool {
	for i := 0; i < kapasitas/2; i++ {
		if arr[i] != arr[kapasitas-1-i] {
			return false
		}
	}
	return true
}

//Buatlah program utamanya untuk menguji tiga subprogram yang telah dibuat tersebut!
func main() {
	var arr arrayInt

	isiArray(&arr, 5)
	fmt.Println("Array sebelum reverse:")
	fmt.Println(arr)
	reverseArray(&arr)
	fmt.Println("Array setelah reverse:")
	fmt.Println(arr)

	if isPalindrom(&arr) {
		fmt.Println("Array memiliki pola palindrom")
	} else {
		fmt.Println("Array tidak memiliki pola palindrom")
	}
}