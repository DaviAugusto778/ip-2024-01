package main

import (
	"fmt"
	"sort"
)

var Tamanho int

func main() {
	fmt.Println("Qual o tamanho da lista?")
	fmt.Scan(&Tamanho)
	fmt.Println("Quais os números da lista?")
	var Lista = make([]float64, Tamanho)
	for a := range Lista {
		fmt.Scan(&Lista[a])
	}
	sort.Float64s(Lista)
	if Tamanho%2 == 0 {
		fmt.Println((Lista[(Tamanho/2)] + Lista[(Tamanho-1)/2]) / 2)
	} else {
		fmt.Println(Lista[(Tamanho)/2])
	}
	fmt.Println(Lista)
}
