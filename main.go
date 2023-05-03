package main

import "fmt"
import "strings"

func main(){
	str := "String For testing some new test"

	fmt.Println("String: ", str)

	fmt.Println("String-uppercase: ", strings.ToUpper(str))
	fmt.Println("String-lowercase: ", strings.ToLower(str))

	// replace test with produce
	fmt.Println("String-replace: ", strings.Replace(str, "test", "produce", -1))
}