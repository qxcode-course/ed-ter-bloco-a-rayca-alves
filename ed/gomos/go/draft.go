package main
import "fmt"

type Gomo struct{
    x, y int
}
func main() {
    var Q int
    dire := "" 
    fmt.Scan(&Q, &dire)

    cobra := make([]Gomo, Q)

    for i := range cobra {
        fmt.Scan(&cobra[i].x, &cobra[i].y)
    }

    for i := Q - 1; i > 0; i--{
        cobra[i] = cobra[i-1]
    }

    switch dire{
    case "L":
        cobra[0].x--
    case "R":
        cobra[0].x++
    case "U":
        cobra[0].y--
    case "D":
        cobra[0].y++
    }

    for _, gomo := range cobra{
        fmt.Printf("%v %v\n", gomo.x, gomo.y)
    }
}
