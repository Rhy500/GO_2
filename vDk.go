package main

import "fmt"

func main() {
	var hur int
	//var hasil bool
	fmt.Scan("%c", &hur)

	if hur == 'a' || hur == 'A' || hur == 'e' || hur == 'E' || hur == 'i' || hur == 'I' || hur == 'u' || hur == 'U' || hur == 'O' || hur == 'o' {
		fmt.Println("vokal")
	} else {
		fmt.Println("konsonal")
	}
}