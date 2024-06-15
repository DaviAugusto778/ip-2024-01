package main
import "fmt"

func exp (x int, y int) int {
    if y == 0 {
        return 1
    }
    return x * exp(x,y - 1)
}

var Num,n int

func main() {
  fmt.Println("Quais os números?")
  fmt.Scan(&Num,&n)
  fmt.Println(exp(Num,n))
}
