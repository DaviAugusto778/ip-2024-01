package main

import (
	"fmt"
	"sort"
)

var Tamanho, Num int
var Lista []int

func main() {
	fmt.Println("Qual o tamanho da lista?")
	fmt.Scan(&Tamanho)
	fmt.Println("Quais os números da lista?")
	for i := 0; i < Tamanho; i++ {
		fmt.Scan(&Num)
		Lista = append(Lista, Num)
	}
	fmt.Println(Lista)
	for a, b := 0, len(Lista)-1; a < b; a, b = a+1, b-1 {
		Lista[a], Lista[b] = Lista[b], Lista[a]
	}
	fmt.Println(Lista)
	sort.Ints(Lista)
	fmt.Println(Lista[len(Lista)-1])
	fmt.Println(Lista[0])
}
