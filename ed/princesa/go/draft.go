package main

import "fmt"

func print(fila []int, espada int) {
	fmt.Print("[ ")
	for i := 0; i < len(fila); i++ {
		if i == espada {
			fmt.Printf("%d> ", fila[i])
		} else {
			fmt.Printf("%d ", fila[i])
		}
	}
	fmt.Println("]")
}

func main() {
	var n, e int
	fmt.Scan(&n, &e)

	fila := make([]int, n)
	for i := 0; i < n; i++ {
		fila[i] = i + 1
	}

	mortos := make(map[int]bool)
	espada := e - 1

	for len(fila) > 1 {
		print(fila, espada)

		morto := (espada + 1) % len(fila)
		mortos[fila[morto]] = true

		nova := []int{}
		for _, p := range fila {
			if !mortos[p] {
				nova = append(nova, p)
			}
		}

        if morto > espada {
			espada = (espada + 1) % len(nova)
		} else {
			espada = espada % len(nova)
		}

		fila = nova
	}

	
	print(fila, espada)
}