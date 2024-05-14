package main

import (
	"fmt"
)

var QuantAlunos, Min, Cont int
var Indices []int

func main() {
	fmt.Println("Quantos alunos e qual a presença minima?")
	fmt.Scan(&QuantAlunos, &Min)
	fmt.Println("Qual a ordem de chegada dos alunos?")
	var Alunos = make([]int, QuantAlunos)
	for i := range Alunos {
		fmt.Scan(&Alunos[i])
		if Alunos[i] <= 0 {
			Cont++
			Indices = append(Indices, i+1)
		}
	}
	if Cont < Min {
		fmt.Println("SIM")
	} else {
		fmt.Println("NÃO")
		for a := len(Indices) - 1; a >= 0; a-- {
			fmt.Println(Indices[a])
		}
	}
}
