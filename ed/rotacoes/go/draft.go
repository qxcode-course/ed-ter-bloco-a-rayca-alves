package main
import "fmt"
func rodar( v[]int, r int) []int{
    n := len(v)
     if n > 0{
        r = r % n
    }
    return  append(v[n-r:], v[:n-r]...)
}
func main() {
    var T, R int
    fmt.Scan(&T, &R)
    v := make([]int, T)

    for i := 0; i < T; i++{
        v[i] = i+1
    }

    fila := rodar(v,R)    
    fmt.Print("[")
    saida := fmt.Sprintf("%v", fila)
    if saida == "[]"{
        fmt.Printf("N ")
        
    }else{
        fmt.Printf(" ")
        fmt.Printf(saida[1 : len(saida) -1])
        fmt.Printf(" ")
    }
    fmt.Println("]")
}
