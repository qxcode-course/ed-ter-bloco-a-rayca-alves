package main
import "fmt"
func main() {
   var C, B int
   fmt.Scan(&C,&B)

   consultas := make([]string, C)
   buscas := make([]string, B)

   

   for i := 0; i < C; i++{
    fmt.Scan(&consultas[i])
   }
   for i := 0; i < B; i++{
    fmt.Scan(&buscas[i])
   }

}
