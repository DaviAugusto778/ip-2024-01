// QUESTÃO 21

package main
import "fmt"
import "sort"

var n, Num int
var Pares = []int{}
var Impares = []int{}

func Imprima (x[]int) {
    for a := range x {
        fmt.Println(x[a])
    }
}

func main() {
  fmt.Println("Quantos números?")
  fmt.Scan(&n)
  fmt.Println("Quais os números?")
  for i := 0; i < n; i++ {
      fmt.Scan(&Num)
      if Num%2 == 0 {
          Pares = append(Pares,Num)
      } else {
          Impares = append(Impares,Num)
      }
  }
  sort.Ints(Pares)
  sort.Ints(Impares)
  Imprima(Pares)
  Imprima(Impares)
}
