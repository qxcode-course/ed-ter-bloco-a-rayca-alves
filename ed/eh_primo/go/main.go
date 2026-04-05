package main

import "fmt"

// x: número que está sendo testado
// div: divisor que está sendo testado
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

func main() {
	var x int
	fmt.Scan(&x)
	if eh_primo(x, 2) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
