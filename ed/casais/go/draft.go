package main

import "fmt"

func main() {
	qtd := 0
	fmt.Scan(&qtd)
	animais := make([]int, qtd)

	solteiros := make(map[int]int)
	for i := range animais{
		fmt.Scan(&animais[i])
	}
	n_pares := 0

	for _, animal := range animais{
		if solteiros[-animal] > 0{
			solteiros[-animal]--
			n_pares++
		}else{
			solteiros[animal]++
		}
	}

	fmt.Println(n_pares)
	/*var N, i int
	fmt.Scan(&N)

	var animais [50]int
	var i_f, im int
	var casais int
	var casado [100]bool

	for i = 0; i < N; i++ {
		fmt.Scan(&animais[i])
		if animais[i] < 0 {
			i_f++
		} else {
			im++
		}
		for j := 1; j < i_f; j++ {
			for x := 1; x < im; i++ {

				if casado[j] == false && animais[x] == (animais[j]*-1) {
					casais++
					casado[j] = true
				}else{
					casado[j] = false
				}
			}
		}

	}

    fmt.Println(casais)
*/
}
