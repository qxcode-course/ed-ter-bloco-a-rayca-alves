package main
import "fmt"
func eh_primo(x int, div int) bool {
	_, _ = x, div
	if x > 1{
		if div == x{
			return true
		}else if x%div == 0{
			return false
		}else{
			return eh_primo(x, div+1)
		}
	}else{
		return false
	}
}

func enesimo(num, n int)int{
	if eh_primo(num, 2) {
		n--
	}
	if n == 0{
		return num
	}
	return enesimo(num+1, n)
}

func main() {
    var num int
    fmt.Scan(&num)

	fmt.Println(enesimo(2, num))
}
