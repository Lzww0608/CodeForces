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

	var n, m, x int
	fmt.Fscan(in, &n, &m)
	vis := make([]bool, n+1)
	for i := 0; i < m*2; i++ {
		fmt.Fscan(in, &x)
		vis[x] = true
	}

	idx := 1
	for vis[idx] {
		idx++
	}
	fmt.Fprintln(out, n-1)
	for i := 1; i <= n; i++ {
		if i != idx {
			fmt.Fprintln(out, idx, i)
		}
	}
}
