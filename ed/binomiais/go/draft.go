package main
import "fmt"
func recur(n, k int)int{
    if k == 0 || k == n{
        return 1
    } 
    return recur(n-1, k - 1) + recur(n-1, k)
}
func main() {
    var n, k int
    fmt.Scan(&n, &k)

    fmt.Println(recur(n, k))

}
