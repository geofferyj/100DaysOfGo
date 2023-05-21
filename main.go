package main

import "fmt"

func main() {

	arr := []string{"a", "b", "c", "d", "e", "f", "g"}

	fmt.Println(deleteElem(arr, "c"))
}

func deleteElem(arr []string, elem string) []string {
	for i, v := range arr {
		if v == elem {
			arr = append(arr[:i], arr[i+1:]...)
		}
	}
	return arr
}