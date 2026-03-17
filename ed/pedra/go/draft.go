package main
import (
    "fmt"
    "math"
)
func main() {
   var N int
   fmt.Scan(&N)

   var vencedor int = -1
   var m_pont int = 1000
   for i := 0; i < N; i++{
    var a, b float64
    fmt.Scan(&a, &b)
    if a >= 10 && b >= 10{
        var pont int = int(math.Abs(a-b))
        
        if pont < m_pont {
            m_pont = pont
            vencedor = i
        }
    }
 }
  if vencedor == -1{
    fmt.Println("sem ganhador")
   }else{
    fmt.Println(vencedor)
   }
}
