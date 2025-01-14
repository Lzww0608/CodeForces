package main

import (
	"bufio"
	"fmt"
	"os"
)

const N int = 1_000_001

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	for fmt.Fscan(in, &t); t > 0; t-- {
		solve(in, out)
	}
}

func solve(in *bufio.Reader, out *bufio.Writer) {
	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	fmt.Fprintln(out, "YES")
	p := 1
	b := make([]bool, N)
	for i := 0; i < n; i++ {
		fmt.Fprint(out, p, " ")
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				x := p + a[j] - a[k]
				if x > 0 && x < N {
					b[x] = true
				}
			}
		}
		for b[p] {
			p++
		}
	}

	fmt.Fprintln(out)
}
