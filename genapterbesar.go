package main

import "fmt"

const N int = 1000

func genapTerbesar(T [N]int, jumlah int) int {
	// Initialize the maximum even number to -1
	maxEven := -1

	// Iterate through the array
	for i := 0; i < jumlah; i++ {
		// Check if the current number is even
		if T[i]%2 == 0 {
			// Update the maximum even number if the current number is greater
			if T[i] > maxEven {
				maxEven = T[i]
			}
		}
	}

	// Return the maximum even number or -1 if no even number is found
	return maxEven
}

func main() {
	// Read the number of elements in the array
	var jumlah int
	fmt.Scanln(&jumlah)

	// Read the elements of the array
	var T [N]int
	for i := 0; i < jumlah; i++ {
		fmt.Scan(&T[i])
	}

	// Print the result
	fmt.Println(genapTerbesar(T, jumlah))
}
