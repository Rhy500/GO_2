package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var totalBelajar, totalMain int

	fmt.Print("Masukkan aktivitas: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	for input != "tidur" {
		switch input {
		case "belajar":
			totalBelajar++
		case "main":
			totalMain++
		}

		fmt.Print("Masukkan aktivitas: ")
		input, _ = reader.ReadString('\n')
		input = strings.TrimSpace(input)
	}

	fmt.Println("Total belajar:", totalBelajar)
	fmt.Println("Total main:", totalMain)
	fmt.Println("Total aktivitas: ", totalBelajar+totalMain)
}
