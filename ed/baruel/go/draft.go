package main
import "fmt"


func main() {
    var qtd_album, qtd_possui int
    fmt.Scan(&qtd_album, &qtd_possui)
    figurinhas := make([]int, qtd_possui)
    for i := range figurinhas{
        fmt.Scan(&figurinhas[i])
    }
    fmt.Println(qtd_album, qtd_possui, figurinhas)
    unicos := make(map[int]bool)
    repetidos := make([]int, 0, qtd_possui)
    for _, figura := range figurinhas{
        if unicos[figura]{
            repetidos = append(repetidos, figura)
        }else{
            unicos[figura] = true
        }
    }

    saida := fmt.Sprintf("%v", repetidos)
    fmt.Println(saida[1 : len(saida)-1])

    faltantes := make([]int, 0, qtd_album)
    for i := 1; i <= qtd_album; i++{
        if !unicos[i]{
            faltantes = append(faltantes, i)
        }
    }

    saida_f := fmt.Sprintf("%v", faltantes)
    fmt.Println(saida[1 : len(saida_f)-1])
    

}
