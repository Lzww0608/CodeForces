package main

import (
	"bufio"
	"fmt"
	"os"
)

const N int = 1_000_000

var m map[int]bool

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	m = make(map[int]bool)
	for k := 2; k*k+k+1 <= N; k++ {
		pre := k * k
		x := k + 1
		for x+pre <= N {
			x += pre
			m[x] = true
			pre *= k
		}
	}

	var t int
	for fmt.Fscan(in, &t); t > 0; t-- {
		solve(in, out)
	}
}

func solve(in *bufio.Reader, out *bufio.Writer) {
	var n int
	fmt.Fscan(in, &n)
	if m[n] {
		fmt.Fprintln(out, "YES")
	} else {
		fmt.Fprintln(out, "NO")
	}
}
