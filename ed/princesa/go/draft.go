package main

import "fmt"

func print(fila []int, pos int) {
	fmt.Print("[ ")
	for i := 0; i < len(fila); i++ {
		if i == pos {
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
	pos := e - 1

	for len(fila) > 1 {
		print(fila, pos)

		morto := (pos + 1) % len(fila)
		mortos[fila[morto]] = true

		nova := []int{}
		for _, p := range fila {
			if !mortos[p] {
				nova = append(nova, p)
			}
		}

        if morto > pos {
			pos = (pos + 1) % len(nova)
		} else {
			pos = pos % len(nova)
		}

		fila = nova
	}

	
	print(fila, pos)
}