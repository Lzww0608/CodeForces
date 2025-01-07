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
	fmt.Fprintln(out, 100_003, 100_003)
	for i := 1; i < n-1; i++ {
		fmt.Fprintln(out, i, i+1, 1)
	}
	fmt.Fprintln(out, n-1, n, 100_003-(n-2))
	k := n
	for i := 1; i < n && k <= m; i++ {
		for j := i + 2; j <= n && k <= m; j++ {
			fmt.Fprintln(out, i, j, 10_000_000)
			k++
		}
	}

}
