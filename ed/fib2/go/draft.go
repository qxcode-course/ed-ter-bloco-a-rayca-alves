package main
import "fmt"
var vet = make(map[int]int)

func recur(n int)int{
   
   if n == 0{
      return 0
   }else if n == 1 || n == 2{
    return 1
   } else if n == 3 || n == 4{
      return 2
   }
   if v, c := vet[n]; c {
      return v
   }
   vet[n] = recur(n - 1) + recur(n - 2) - recur(n - 4)
   return vet[n]
}
func main() {
   var n int
   fmt.Scan(&n)

   fmt.Println(recur(n))
}
