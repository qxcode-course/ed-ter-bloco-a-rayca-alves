package main

import "fmt"

func main() {
	var N, i int
	fmt.Scan(&N)

	var animais [50]int
	var f [100]int
	var i_f, im int
	var m [100]int
	var pares int
	casado := false

	for i = 0; i < N; i++ {
		fmt.Scan(&animais[i])
		if animais[i] < 0 {
			i_f++
			f[i] = animais[i]
		} else {
			im++
			m[i] = animais[i]
		}
		for j := 1; j < i_f; j++ {
			for x := 1; x < im; i++ {

				if casado == false && m[x] == (f[j]*-1) {
					pares++
					casado = true
				}
			}
		}

	}

    fmt.Println(pares)

}
