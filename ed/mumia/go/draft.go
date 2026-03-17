package main
import "fmt"
func main() {
   var nome string
   fmt.Scan(&nome)

   var idade int
   fmt.Scan(&idade)

   fmt.Print(nome, " eh ")
   if idade < 12{
    fmt.Println("crianca")
   }else if idade < 18{
    fmt.Println("jovem")
   }else if idade < 65{
    fmt.Println("adulto")
   }else if idade < 1000{
    fmt.Println("idoso")
   }else{
    fmt.Println("mumia")
   }
}
