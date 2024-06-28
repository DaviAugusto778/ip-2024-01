package main
import "fmt"

func troca(x[]int,a,b int) {
    x[a],x[b] = x[b],x[a]
}

func trocaOpostosSeMenor(x[]int) {
    for a,b := 0, len(x) - 1; b > a ; a,b = a + 1,b - 1 {
        if x[a] < x[b] {
            troca(x,a,b)
        }
    }
}


var n,casos int

func main() {
  fmt.Println("Quantos casos de teste?")
  fmt.Scan(&casos)
  for k := 0; k < casos; k++ {
  fmt.Println("Qual o tamanho do vetor")
  fmt.Scan(&n)
  var Vetor = make([]int,n)
  fmt.Println("Quais os numeros?")
  for i := range Vetor{
      fmt.Scan(&Vetor[i])
  }
      fmt.Println(Vetor)
      trocaOpostosSeMenor(Vetor)
      fmt.Println(Vetor)
  }
}
