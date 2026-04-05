package main
import "fmt"
func recur(n, k int)int{
   if n == 1 || n == 2{
    return 1
   }
   return recur(n - 1, k) + k* recur(n - 2, k)
}
func main() {
   var n, k int
   fmt.Scan(&n, &k)

   fmt.Println(recur(n,k))
}
