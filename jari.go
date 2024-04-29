package main

import (
	"fmt"
	"math"
)

const pi float64 = 3.14

type Bola struct {
	Nama     string
	Jarijari int
	Volume   float64
}

func main() {
	// membuat slince/masukan  untuk menyimpan data bola
	bola := make([]Bola, 5)

	//membaca imput data bola
	for i := 0; i < 5; i++ {
		var nama string
		var jarijari int
		fmt.Scan(&nama, &jarijari)

		//menghitung volume bola
		volume := jaribola(jarijari)

		// menyimpan informasi bola ke dalam slice
		bola[i] = Bola{Nama: nama, Jarijari: jarijari, Volume: volume}
	}
	// mencari bola dengan volume terbesar
	bolaTerbesar := bola[0]
	for _, bola := range bola {
		if bola.Volume > bolaTerbesar.Volume {
			bolaTerbesar = bola
		}
	}
	fmt.Printf("%s %.2f\n", bolaTerbesar.Nama, bolaTerbesar.Volume)
}
func jaribola(jarijari int) float64 {
	return (4.0 / 3.0) * pi * math.Pow(float64(jarijari), 3)

}
