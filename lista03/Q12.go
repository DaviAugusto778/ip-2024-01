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
	sort.Ints(Lista)
	for a := 0; a < len(Lista); a++ {
		fmt.Println(Lista[a])
	}
}
