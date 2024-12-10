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

	var n, m int
	fmt.Fscan(in, &n, &m)
	a := make([][]int, n)
	b := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int, m)
		b[i] = make([]int, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Fscan(in, &a[i][j])
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Fscan(in, &b[i][j])
		}
	}

	cnt := make([]map[int]int, n+m+1)
	for i := range cnt {
		cnt[i] = make(map[int]int)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			cnt[i+j][a[i][j]]++
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			x := b[i][j]
			if cnt[i+j][x]--; cnt[i+j][x] < 0 {
				fmt.Fprintln(out, "NO")
				return
			}
		}
	}
	fmt.Fprintln(out, "YES")
}
