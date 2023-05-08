package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your score: ")
	text, _ := reader.ReadString('\n')

	score, _ := strconv.Atoi(strings.TrimSpace(text))
	var grade string

	switch {
	case score >= 80:
		grade = "A"
	case score >= 70:
		grade = "B"
	case score >= 60:
		grade = "C"
	case score >= 50:
		grade = "D"
	default:
		grade = "F"
	}

	fmt.Println("Your grade is", grade)

}
