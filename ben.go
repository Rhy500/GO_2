package main

import "fmt"

func main() {
	var M, N int

	fmt.Printf("Berapa kekuatan awal gerbang? ")
	fmt.Scan(&M)
	fmt.Printf("Berapa daya rusak balok? ")
	fmt.Scan(&N)

	fmt.Println("Kekuatan awal gerbang ", M)
	dobrakPintu(M, N, 0)
}

func dobrakPintu(M, N, O int) {
	if M >= 0 {
		fmt.Println("dobrak, gerbang jebol")
		fmt.Println("Pasukan berusaha sebanyak", O, "kali")
		return
	}

	fmt.Println("dobrak, sisa kekuatan", M-N)
	dobrakPintu(M-N, N, O+1)

}
