package main

import (
	"fmt"
)

var Tamanho, Num, NumMax, Index int
var Lista [101]int

func main() {
	fmt.Println("Quantos números?")
	fmt.Scan(&Tamanho)
	fmt.Println("Quais os números?")
	for i := 0; i < Tamanho; i++ {
		fmt.Scan(&Num)
		Lista[Num]++
	}
	for i := range Lista {
		if Lista[i] > NumMax {
			NumMax = Lista[i]
			Index = i
		}
	}
	fmt.Println("-------------------")
	fmt.Print(Index, "\n", NumMax)
}
