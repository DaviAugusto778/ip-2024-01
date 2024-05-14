// Questão 8

package main
import "fmt"
import "strconv"

var Num int64

func main() {
  fmt.Println("Diga os números para transformar em binário:")
  for {
      fmt.Scan(&Num)
      fmt.Println(strconv.FormatInt(Num,2))
  }
}
