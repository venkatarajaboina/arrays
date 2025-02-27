package main

import "fmt"

func main() {
	numbers := [3]int{10, 20, 30} // Declare and initialize an array

	fmt.Println("Array:", numbers)    // Print array
	fmt.Println("First:", numbers[0]) // Access first element

	// Loop through array
	for i, v := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}
}
