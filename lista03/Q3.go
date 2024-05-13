// QUESTÃO 3

package main

import (
	"fmt"
	"slices"
)

var Tamanho int

func main() {
	fmt.Println("Qual o tamanho da lista?")
	fmt.Scan(&Tamanho)
	fmt.Println("Quais os números da lista?")
	var Lista = make([]int, Tamanho)
	for a := range Lista {
		fmt.Scan(&Lista[a])
	}
	slices.Reverse(Lista)
	fmt.Println(Lista)
}
