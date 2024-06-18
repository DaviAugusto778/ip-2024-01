package main

import "fmt"

func inv(x []int, y int) {
	if y == 0 {
		fmt.Println(x[y])
	} else {
		fmt.Println(x[y])
		inv(x, y-1)
	}
}

var n int

func main() {
	fmt.Println("Quantos números")
	fmt.Scan(&n)
	fmt.Println("Quais os números?")
	var Lista = make([]int, n)
	for i := range Lista {
		fmt.Scan(&Lista[i])
	}
	fmt.Println("A lista normal:", Lista)
	inv(Lista, n-1)
}
