package main
import "fmt"

func Crescente(x[]int) {
    for a := len(x); a > 0; a-- {
        for b := 1; b < a; b++ {
            if x[b-1] > x[b] {
                x[b-1],x[b] = x[b],x[b-1]
            }
        }
    }
}

var n int

func main() {
  fmt.Println("Qual o tamanho do vetor")
  fmt.Scan(&n)
  var Vetor = make([]int,n)
  fmt.Println("Quais os numero?")
  for i:= range Vetor {
      fmt.Scan(&Vetor[i])
  }
  fmt.Println(Vetor)
  Crescente(Vetor)
  fmt.Println(Vetor)
}
