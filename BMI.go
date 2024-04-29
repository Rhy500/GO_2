package main

import "fmt"

// Definisikan tipe bentukan untuk review series
type dataTubuh struct {
	tinggi float64
	berat  float64
	bmi    float64
}

func main() {
	var dt dataTubuh
	fmt.Scan(&dt.tinggi, &dt.berat)
	// Calculate BMI
	hasil(&dt)

	// Output BMI and category
	fmt.Printf("BMI: %.2f\n", dt.bmi)
	fmt.Println("Kategori:", kategori(dt.bmi))
}
func hasil(dt *dataTubuh) {
	dt.bmi = (dt.berat / (dt.tinggi * dt.tinggi))
}
func kategori(bmi float64) string {
	if bmi < 18.5 {
		return ("Underweight")
	} else if bmi >= 18.5 && bmi < 24.9 {
		return ("Normal weight")
	} else if bmi >= 25 && bmi < 29.9 {
		return ("Overweight")
	} else {
		return ("Obesity")
	}
}
