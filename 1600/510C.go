package main

import (
	"bufio"
	"fmt"
	"os"
)

const N int = 26

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)
	a := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	deg := make([]int, N)
	g := make([][]int, N)
	for i := 0; i < n; i++ {
	next:
		for j := i + 1; j < n; j++ {
			for k := 0; k < len(a[i]) && k < len(a[j]); k++ {
				if a[i][k] != a[j][k] {
					x, y := int(a[i][k]-'a'), int(a[j][k]-'a')
					if x == y {
						continue
					}
					if deg[y]&(1<<x) == 0 {
						deg[y] |= 1 << x
						g[x] = append(g[x], y)
					}
					continue next
				}
			}
			if len(a[i]) > len(a[j]) {
				fmt.Fprintln(out, "Impossible")
				return
			}
		}
	}

	vis := make([]bool, N)
	q := make([]int, 0, N)
	for i, x := range deg {
		if x == 0 {
			q = append(q, i)
			vis[i] = true
		}
	}

	ans := make([]byte, 0, N)
	for len(q) > 0 {
		x := q[0]
		ans = append(ans, byte('a'+x))
		q = q[1:]
		for _, y := range g[x] {
			deg[y] &^= 1 << x
			if deg[y] == 0 && !vis[y] {
				vis[y] = true
				q = append(q, y)
			}
		}
	}

	if cap(q) != 0 {
		fmt.Fprintln(out, "Impossible")
		return
	}
	fmt.Fprintln(out, string(ans))
}
