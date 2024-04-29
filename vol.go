package main

import "fmt"

func vowel(kar byte) bool {
	return kar == 'a' || kar == 'i' || kar == 'u' || kar == 'e' || kar == 'o' || kar == 'A' || kar == 'I' || kar == 'U' || kar == 'E' || kar == 'O'

}
func main() {
	var kar byte
	var vow int
	fmt.Scanf("%c", &kar)
	for kar != ',' {
		if vowel(kar) {
			vow = vow + 1
		}
		fmt.Scanf("%c", &kar)
	}
	fmt.Println("Jumlahkan vokal adalah", vow)
}
