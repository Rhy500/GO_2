package main

import "fmt"

func jumlahA(x string, index int) int {
	/* I.S Terdefinisi nilai string x dan bilangan bulat index
	F.S mengembalikan jumlah huruf 'a' dari string x */
	if index < len(x) {
		//if strings.ToLower(string(x[index])) == "a" {
		if string(x[index]) == "a" {
			return 1 + jumlahA(x, index+1)
		} else {
			return jumlahA(x, index+1)
		}
	} else {
		return 0
	}
}

func main() {
	var x string
	var result int
	fmt.Scanln(&x)
	result = jumlahA(x, 0)
	fmt.Println(result)
}
