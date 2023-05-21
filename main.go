package main

import "fmt"

func main(){
	s := [7]string{"mango", "banana", "apple", "kiwi", "pineapple", "orange", "grape"}

	for index, value := range s {
		s[len(s) - index-1] = value
	}

	fmt.Println(s)
}