package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t, n int
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &n)
		a := make([]int, n*2)
		for i := range a {
			fmt.Fscan(in, &a[i])
		}
		i, j := 0, 0
		k := 0
		for i < n*2 && k < n*2-2 {
			x, y := -1, -1
			for i < n*2 && a[i]&1 == 0 {
				i++
			}
			x = i
			i++
			for i < n*2 && a[i]&1 == 0 {
				i++
			}
			y = i
			i++
			if x >= 0 && x < n*2 && y >= 0 && y < n*2 {
				k += 2
				fmt.Fprintln(out, x+1, y+1)
			}

		}

		for j < n*2 && k < n*2-2 {
			x, y := -1, -1
			for j < n*2 && a[j]&1 == 1 {
				j++
			}
			x = j
			j++
			for j < n*2 && a[j]&1 == 1 {
				j++
			}
			y = j
			j++
			if x >= 0 && x < n*2 && y >= 0 && y < n*2 {
				k += 2
				fmt.Fprintln(out, x+1, y+1)
			}

		}
	}
}
