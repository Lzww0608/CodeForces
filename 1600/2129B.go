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

	var t int
	for fmt.Fscan(in, &t); t > 0; t-- {
		solve(in, out)
	}
}

func solve(in *bufio.Reader, out *bufio.Writer) {
	var n int
	fmt.Fscan(in, &n)
	p := make([]int, n)
	for i := range p {
		fmt.Fscan(in, &p[i])
	}

	ans := 0
	for i := range n {
		a, b := 0, 0
		for j := i - 1; j >= 0; j-- {
			if p[j] > p[i] {
				a++
			}
		}

		for j := i + 1; j < n; j++ {
			if p[j] > p[i] {
				b++
			}
		}

		ans += min(a, b)
	}

	fmt.Fprintln(out, ans)
}
