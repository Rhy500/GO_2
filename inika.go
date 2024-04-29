package main

import "fmt"

const N = 7 // assuming this is the length of the array T

func hereThere(T [N]int, total, inikah int) int {
	// binary search algorithm
	low := 0
	high := total - 1

	for low <= high {
		mid := low + (high-low)/2

		if T[mid] == inikah {
			return inikah
		} else if T[mid] < inikah {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	// if inikah is not found, return the next largest value or -1 if all values are smaller
	if low < N && T[low] > inikah {
		return T[low]
	} else {
		return -1
	}
}

func main() {
	// sample input
	T := [N]int{2, 3, 5, 7, 11, 13, 17}
	total := N
	// ubah nilai inikah ini untuk nilai inikah lainnya
	inikah := 21

	// call the hereThere function
	retValue := hereThere(T, total, inikah)

	// print the result
	fmt.Printf("The result is: %d\n", retValue)
}
