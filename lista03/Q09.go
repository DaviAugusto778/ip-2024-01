package main

import (
	"fmt"
	"math"
)

var n int
var Num float64
var Lista []float64

func main(){
    fmt.Println("Quantos pontos?")
	fmt.Scan(&n)
	fmt.Println("Quais os pontos?")
	for i:=0;i < n;i++{
		for a:=0;a < 3;a ++{
			fmt.Scan(&Num)
			Lista = append(Lista, Num)
		}
	}
	for k := 0; k < n * 3 - 3;k += 3{
		fmt.Printf("%.2f\n",math.Sqrt(float64((Lista[k]-Lista[k+3])*(Lista[k]-Lista[k+3])+(Lista[k+1]-Lista[k+4])*(Lista[k+1]-Lista[k+4])+(Lista[k+2]-Lista[k+5])*(Lista[k+2]-Lista[k+5]))))
	}
}
