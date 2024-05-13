// QUESTÃO 02

package main
import "fmt"

var Tamanho,Num,Res int

func main() {
  fmt.Print("Qual o tamanho da lista? ")
  for Tamanho < 1 || Tamanho > 1000 {
  fmt.Scan(&Tamanho)
  if Tamanho >= 1 && Tamanho <= 1000 {
  fmt.Println("Quais os números da lista?")
  var Lista = make([]int,Tamanho)
  for i := range Lista {
      fmt.Scan(&Lista[i]) }
  fmt.Print("Qual o número a ser comparado? ")
  fmt.Scan(&Num)
  for a := 0; a < len(Lista); a++ {
      if Num <= Lista[a] {
          Res++
      }
      }
      fmt.Println(Res)
  }
  }
}
