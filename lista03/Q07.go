package main

import "fmt"

func Contagem(Lista []int, Tamanho int) int {
	var Cont int
	for a := range Lista {
		if Tamanho == Lista[a] {
			Cont++
		}
	}
	return Cont
}

var Tamanho, Num, NumMax, Soma int
var Lista = []int{}

func main() {
	for {
		fmt.Println("Qual o tamanho da lista?")
		fmt.Scan(&Tamanho)
		if Tamanho == 0 || Tamanho >= 1000 {
			break
		}
		fmt.Println("Quais os números da lista?")
		for a := 0; a < Tamanho; a++ {
			fmt.Scan(&Num)
			Lista = append(Lista, Num)
			if Num > NumMax {
				NumMax = Num
			}
		}
		for a := 0; a <= NumMax; a++ {
			Soma += Contagem(Lista, a)
			fmt.Printf("(%v) %v\n", a, Soma)
		}
		Soma = 0
	}
}
