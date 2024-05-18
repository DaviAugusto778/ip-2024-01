package main

import (
	"fmt"
)

var n, Tamanho, Num, Index int

func main() {
	fmt.Println("Quantos números?")
	fmt.Scan(&Tamanho)
	fmt.Println("Quais os números?")
	var Lista = make([]int, Tamanho+1000)
	for i := 0; i < Tamanho; i++ {
		fmt.Scan(&Num)
		Lista[Num]++
	}
	for a := range Lista {
		if Lista[a] == 1 {
			Index++
		}
	}
	fmt.Println(Index)
}
