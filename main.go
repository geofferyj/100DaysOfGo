package main

import "fmt"

func main() {
	num1 := 100
	num2 := 20
	num3 := 30


	//find the largest number

	if num1 >= num2 && num1 >= num3 {
		fmt.Println("num1 is the largest number")
	} else if num2 >= num1 && num2 >= num3 {
		fmt.Println("num2 is the largest number")
	} else {
		fmt.Println("num3 is the largest number")
	}
}