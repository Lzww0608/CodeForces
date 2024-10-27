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
	var n, k int
	fmt.Fscan(in, &n, &k)
	mx := n * n / 2
	if k&1 == 1 || k > mx {
		fmt.Fprintln(out, "NO")
		return
	}
	fmt.Fprintln(out, "YES")
	vis := make([]bool, n)
	p := make([]int, n)
	for i := range p {
		p[i] = -1
	}

	k >>= 1
	for i := 0; i < n-1-i && k > 0; i++ {
		t := min(n-1-i, i+k)
		vis[t] = true
		p[i] = t
		k -= t - i
	}

	for i, j := 0, 0; i < n; i++ {
		if p[i] == -1 {
			for vis[j] {
				j++
			}
			p[i] = j
			j++
		}

		fmt.Fprint(out, p[i]+1, " ")
	}
	fmt.Fprintln(out)
}
