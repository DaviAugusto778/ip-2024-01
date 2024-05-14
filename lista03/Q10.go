package main

import (
	"fmt"
)

var Tamanho, Num, NumMax, Ref, Cont, Index int
var Lista []int

func main() {
	fmt.Println("Qual o tamanho da lista?")
	fmt.Scan(&Tamanho)
	fmt.Println("Quais os números da lista?")
	for i := 0; i < Tamanho; i++ {
		fmt.Scan(&Num)
		Lista = append(Lista, Num)
		if Num > NumMax {
			NumMax = Num
			Index = i
		}
	}
	for a := 0; a < len(Lista); a++ {
		if Lista[a] == Lista[len(Lista)-1] {
			Cont++
		}
	}
	fmt.Printf("Nota %v, %v vezes\nNota %v, indice %v", Lista[len(Lista)-1], Cont, NumMax, Index)
}
