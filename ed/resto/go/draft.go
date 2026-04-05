package main
import "fmt"
func recur(num int){

    if num == 0{
       return 
    }
    n := num / 2
    x := num % 2

    recur(n)
    fmt.Println(n, x)
}
func main() {
    var num int
    fmt.Scan(&num)

    recur(num)
}
