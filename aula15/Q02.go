package main
import "fmt"

var i int

func soma (x[]int,y int) int {
    if y < 0 {
        return 0
    }
    return x[y] + soma(x, y - 1)
}

var n int

func main() {
  fmt.Println("Quantos números")
  fmt.Scan(&n)
  var Lista = make([]int,n)
  for i := range Lista{
      fmt.Scan(&Lista[i])
  }
  fmt.Println(soma(Lista,n - 1))
}
