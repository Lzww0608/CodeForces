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
	p := make([]int, n+1)
	if n&1 == 1 {
		for i := 1; i <= n-2; i++ {
			p[i] = i + 1
		}
		p[n], p[n-1], p[n-2] = n, n-1, 1
	} else {
		k := 1
		for (1 << k) <= n {
			k++
		}
		k--
		vis := make([]bool, n+1)

		t := n
		for i := k; i > 0; i-- {
			p[t] = 1 << i
			p[t-1] = (1 << i) - 1
			vis[p[t]], vis[p[t-1]] = true, true
			t -= 2
		}

		for i := 1; i <= n; i++ {
			if !vis[i] {
				p[t] = i
				t--
			}
		}
	}
	ans := 0
	for i := 1; i <= n; i++ {
		if i&1 == 1 {
			ans &= p[i]
		} else {
			ans |= p[i]
		}
	}
	fmt.Fprintln(out, ans)
	for i := 1; i <= n; i++ {
		fmt.Fprint(out, p[i], " ")
	}
	fmt.Fprintln(out)
}
