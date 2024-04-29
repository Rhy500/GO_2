package main

import "fmt"

// Definisikan tipe bentukan untuk review series
type dataTubuh struct {
	tinggi float64
	berat  float64
}

func main() {
	var tinggi, berat float64
	fmt.Scan(&tinggi, &berat)
	BMI := berat / (tinggi * tinggi)

}
