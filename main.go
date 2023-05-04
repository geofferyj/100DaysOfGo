package main

import (
	"fmt"
	"bufio"
	"strings"
	"os"
)

func main() {

	fmt.Print("enter a sentence: ")
	scanner := bufio.NewReader(os.Stdin)

	input, _ := scanner.ReadString('\n')

	wordCount := len(strings.Split(input, " "))

	fmt.Printf("there are %d words in the sentence\n", wordCount)


}