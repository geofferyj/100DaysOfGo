package main

import (
	"fmt"
)

func main(){
Fibo(10)
}

func Fibo(n int) {
	f0 := 0
	f1 := 1
	f2 := 1

	for i :=0; i<=n; i++ {
		fmt.Println(f0)
		f0 = f1
		f1 = f2
		f2 = f0 + f1
	}
}

