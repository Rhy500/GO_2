package main

import "fmt"

const N = 1000

func lebihBesar(T [N]int, jumlah, target int) int {
	// Initialize the total to 0
	total := 0

	// Iterate through the array
	for i := 0; i < jumlah; i++ {
		// Check if the current number is greater than the target
		if T[i] > target {
			// Add the current number to the total
			total += T[i]
		}
	}

	// Return the total
	return total
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

	// Read the target value
	var target int
	fmt.Scan(&target)

	// Print the result
	fmt.Println(lebihBesar(T, jumlah, target))
}
