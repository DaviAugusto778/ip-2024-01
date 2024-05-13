// QUESTÃO 01

package main
import "fmt"

var Num,Tamanho,Quantidade int

func contains(Lista []int, n int) string {
	for _, v := range Lista {
		if v == n {
			return "Achei"
		}
	}

	return "Não Achei"
}

func main() {
  fmt.Println("Qual o tamanho da lista?")
  fmt.Scan(&Tamanho)
  fmt.Println("Quais os números da lista?")
  var Lista = make([]int,Tamanho)
  for i := range Lista {
      fmt.Scan(&Lista[i])
  }
	fmt.Println("Quantos números a serem escaneados?")
	fmt.Scan(&Quantidade)
	fmt.Println("Quais os números?")
	for a := 0; a < Quantidade; a++ {
	fmt.Scan(&Num)
	defer fmt.Println(Num,contains(Lista, Num))
	}
}
