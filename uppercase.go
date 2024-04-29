package main

import "fmt"

func lowToUpper(char byte) byte {
	return 'A' + (char - 'a')
}

func main() {
	var char byte
	_, err := fmt.Scanf("%c", &char)
	if err != nil {
		//fmt.Println("Invalid char (must be a single lowercase letter)")
		return
	}
	if char >= 'a' && char <= 'z' {
		output := lowToUpper(char)
		fmt.Printf("%c is %c.\n", output)
	}
}
