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
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}
	b := make([]int, n+1)
	c := make([]int, n+1)
	vis := make([]int, n+1)
	for i := 1; i <= n; i++ {
		if a[i] != a[i-1] {
			vis[a[i]] = a[i]
			b[i] = a[i]
			c[i] = a[i]
		}
	}

	j := 1
	for i := 1; i <= n; i++ {
		if b[i] == 0 {
			for vis[j] != 0 {
				j++
			}
			b[i] = j
			j++
		}
	}

	j = 1
	for i := 1; i <= n; i++ {
		if c[i] == 0 {
			t := j
			for vis[j] != 0 {
				j = vis[j] - 1
			}
			vis[j] = j
			vis[t] = j
			c[i] = j
		} else {
			j = c[i] - 1
		}
	}
	for i := 1; i <= n; i++ {
		fmt.Fprint(out, b[i], " ")
	}
	fmt.Fprintln(out)
	for i := 1; i <= n; i++ {
		fmt.Fprint(out, c[i], " ")
	}
	fmt.Fprintln(out)
}
