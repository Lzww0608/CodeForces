package main

import (
	"bufio"
	"fmt"
	"os"
)

const MOD int = 1_000_000_007

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
	var n, x int
	fmt.Fscan(in, &n)
	p := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &x)
		p[x] = i
	}

	l, r := p[0], p[0]
	ans := 1
	for i, x := range p[1:] {
		if x < l {
			l = x
		} else if x > r {
			r = x
		} else {
			ans = ans * (r - l - i) % MOD
		}
	}
	fmt.Fprintln(out, ans)
}
