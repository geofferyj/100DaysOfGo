package main

import "fmt"

func main() {
	// Input the three numbers
	var a, b, c int
	fmt.Println("Enter three integers separated by space:")
	fmt.Scanln(&a, &b, &c)

	// Determine the largest number
	largest := a
	if b > largest {
		largest = b
	}
	if c > largest {
		largest = c
	}

	// Print the result
	fmt.Printf("The largest number is: %d\n", largest)
}