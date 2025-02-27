package main

import "fmt"

func main() {
	num := 10   // Declare an integer variable
	ptr := &num // Create a pointer to 'num'

	fmt.Println("Value of num:", num)
	fmt.Println("Pointer address:", ptr)      // Prints memory address
	fmt.Println("Value using pointer:", *ptr) // Dereferencing pointer

	*ptr = 20 // Change value using pointer
	fmt.Println("Updated value of num:", num)
}
