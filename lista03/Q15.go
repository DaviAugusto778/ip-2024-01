package main

import (
	"fmt"
	"math"
)

var Quantidade int

func main() {
	fmt.Println("Quantos casos de teste?")
	fmt.Scan(&Quantidade)
	for b := 0; b < Quantidade; b++ {
		var n, Min, Index = 0, 2000, 0
		fmt.Println("Quantos números?")
		fmt.Scan(&n)
		var Lista = make([]int, n)
		fmt.Println("Quais os números?")
		for a := range Lista {
			fmt.Scan(&Lista[a])
		}
		for i := 0; i < n-1; i++ {
			for k := i + 1; k < n; k++ {
				if math.Abs(float64(Lista[i])-float64(Lista[k])) < float64(Min) {
					Min = int(math.Abs(float64(Lista[i]) - float64(Lista[k])))
				}
				Index++
			}
		}
		fmt.Println(Min, Index)
	}
}
