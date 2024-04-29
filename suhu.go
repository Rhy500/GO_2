package main

import "fmt"

func tofahrenheit(cc int) int {
	return ((cc * 9) / 5) + 32
}

func toreamur(cc int) int {
	return (cc * 4) / 5
}

func tokelvin(cc int) int {
	return cc + 273
}

//input 0 100 5

func main() {
	var x, y, z int
	fmt.Scan(&x, &y, &z)
	fmt.Println("Celcius ", "Reamur ", "Fahrenheit ", "kelvin")
	for x <= y {
		fmt.Printf("%.2f %.2f %.2f %.2f\n", float64(x), float64(toreamur(x)), float64(tofahrenheit(x)), float64(tokelvin(x)))
		x += z
	}
}
