package main

import "fmt"

func Binario(n int) {
	if n/2 != 0 {
		Binario(n / 2)
	}
	fmt.Print(n % 2)
}

var n int

func main() {
	fmt.Println("Qual o numero")
	fmt.Scan(&n)
	Binario(n)
}
