package main

import "fmt"

func main(){
	numbers := []int{1,3,4,5,6,5,4,5,6}
	sum := 0

	for index, number := range numbers {
		fmt.Println(index, number)
		sum += number
	}

	fmt.Println("Sum: ", sum)
}