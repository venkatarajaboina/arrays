package main

import "fmt"

func main() {
	fmt.Println("Start") // This executes first

	defer fmt.Println("Deferred Statement") // This runs at the end

	fmt.Println("End") // This executes before the deferred statement
}
