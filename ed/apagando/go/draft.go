package main
import "fmt"
func main() {
    var N, M int
    fmt.Scan(&N)
    fila := make([]int, N)
    for i := 0; i < N; i++{
        fmt.Scan(&fila[i])
    } 

    fmt.Scan(&M)
    apagado := make(map[int]bool)
    sair := make([]int, M)
    for i := 0; i < M; i++{
        fmt.Scan(&sair[i])
        apagado[sair[i]] = true
    } 

    novaFila := []int{}
    for _, f := range fila{
        if !apagado[f]{
            novaFila = append(novaFila, f)
        }
    }
    fila = novaFila

    saida := fmt.Sprintf("%v", fila)
    if saida == "[]"{
        fmt.Printf("N ")
        
    }else{
        fmt.Printf(saida[1 : len(saida) -1])
        fmt.Println(" ")
    }
}

