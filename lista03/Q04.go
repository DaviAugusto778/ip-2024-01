// QUESTÃO 4

package main
import "fmt"

var Tamanho int

func main() {
fmt.Println("Qual o tamanho da lista?")
  fmt.Scan(&Tamanho)
  fmt.Println("Quais os números da lista?")
  var Lista = make([]int,Tamanho)
  for i := range Lista {
      fmt.Scan(&Lista[i])
  }
  fmt.Print(Lista)
}
