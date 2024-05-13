package main

import (
	"fmt"
)

var Tamanho, Num, NumMax, Index int

func main() {
	for {
		fmt.Println("Qual o tamanho da lista?")
		fmt.Scan(&Tamanho)
		if Tamanho == 0 || Tamanho >= 1000 {
			return
		} else {
			fmt.Println("Quais os números da lista?")
			for i := 0; i < Tamanho; i++ {
				fmt.Scan(&Num)
				if Num > NumMax {
					NumMax = Num
					Index = i
				}
			}
			fmt.Println(Index, NumMax)
			NumMax = 0
			Index = 0
		}
	}
}
