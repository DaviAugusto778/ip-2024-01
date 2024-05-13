package main

import "fmt"

var Tamanho, Num, Soma int

func main() {
	fmt.Println("Qual o tamanho da lista?")
	fmt.Scan(&Tamanho)
	for i := 0; i < Tamanho; i++ {
		fmt.Scan(&Num)
		Soma += Num
	}
	fmt.Print(Soma)
}
