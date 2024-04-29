package main

import "fmt"

func main() {
	var tinggi, radius, volume float64
	fmt.Scan(&radius, &tinggi)
	volume = volumeSilinder(radius, tinggi)
	fmt.Println(volume)
}
func volumeSilinder(radius, tinggi float64) float64 {
	const pi float64 = 3.14
	var volume float64
	volume = pi * radius * radius * tinggi
	return volume
}
