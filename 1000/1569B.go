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
		var s string
		id := []int{}
		fmt.Fscan(in, &s)
		for i, c := range s {
			if c == '2' {
				id = append(id, i)
			}
		}
		k := len(id)
		if k == 1 || k == 2 {
			fmt.Fprintln(out, "NO")
			continue
		}
		fmt.Fprintln(out, "YES")
		a := make([][]byte, n)
		for i := range a {
			a[i] = make([]byte, n)
			for j := range a[i] {
				if i == j {
					a[i][j] = 'X'
				} else {
					a[i][j] = '='
				}

			}
		}
		for i := 0; i < k; i++ {
			x := id[i]
			y := id[(i+1)%k]
			a[x][y] = '+'
			a[y][x] = '-'
		}
		for i := range a {
			for j := range a[i] {
				fmt.Fprint(out, string(a[i][j]))
			}
			fmt.Fprintln(out)
		}
	}
}
