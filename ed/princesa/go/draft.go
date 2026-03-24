package main
import "fmt"
func main() {
    var n, e int
    fmt.Scan(&n, &e)

    fila := make([]int, 1, n)
    mortos := make(map[int]int)

    for matar := range fila{
        if fila[matar] == e{
            mortos[matar] = fila[matar+1]
        }
    }
}
