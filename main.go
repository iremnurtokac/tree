package main

import (
	"fmt"
)

func write(comp string, count int) {
	for i := 0; i < count; i++ {
		fmt.Print(comp)
	}
}

func main() {
	n := 4

	for k := 0; k < n; k++ {

		for i := k; i < n+2*k; i++ {

			spacecount := (n - i) + (2*n - 2)
			xcount := (2*i + 1)
			write(" ", spacecount)
			write("x", xcount)
			fmt.Println()
		}
	}
}
