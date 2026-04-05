package main
import "fmt"
func recur(n int)int{
    if n == 1 || n == 2{
        return 1
    }else if n == 3{
        return 2
    }else{
        return recur(n-1) + recur(n - 3)
    } 
}
func main() {
    var n int
    fmt.Scan(&n)

    fmt.Println(recur(n))
}
