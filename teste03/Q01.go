
package main
import "fmt"

var palavra string
var tamanho int

func inv(x string) string {
    letras := []rune(x)
    for a,b := 0, len(letras) -1 ; a < b ; a,b = a+1, b-1{
        letras [a], letras[b] = letras[b],letras[a]
    }
    return string(letras)
}

func main() {
  fmt.Println("Qual a frase?")
  var frase = []string{}
  for {
      fmt.Scan(&palavra)
      palavra = inv(palavra)
      if palavra == "#" {
          break
      } else {
          frase = append(frase,palavra)
      }
  }
  for i := len(frase) - 1 ; i >= 0; i-- {
      fmt.Print(frase[i]," ")
  }
}
