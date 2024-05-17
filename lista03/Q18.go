// Questão 18

package main
import "fmt"

var n,Num int

func main() {
  fmt.Println("Quantos CPFs?")
  fmt.Scan(&n)
  for i := 0 ; i < n ; i++ {
  var CPF = make([]int,11)
  var Multi1,Multi2,Soma1,Soma2 int
      fmt.Println("Qual os digitos do CPF?")
      for a := range CPF {
          fmt.Scan(&CPF[a])
      }
      for k := 0 ; k < 9; k++ {
          Multi1 = CPF[k] * (k + 1)
          Multi2 = CPF[8-k] * (k + 1)
          Soma1 += Multi1
          Soma2 += Multi2
      }
      if Soma1%11 == CPF[9] && Soma2%11 == CPF[10]{
          fmt.Println("CPF valido")
      } else {
          fmt.Println("CPF invalido")
      }
  }
}
