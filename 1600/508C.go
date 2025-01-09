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

	var m, t, r int
	fmt.Fscan(in, &m, &t, &r)
	w := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &w[i])
	}

	if t < r {
		fmt.Fprintln(out, -1)
		return
	}

	ans := 0
	a := 0
	vis := make(map[int]bool)
	for _, x := range w {
		c := 0
		for i := 0; i < t; i++ {
			if vis[x-i] {
				c++
			}
		}
		if c >= r {
			continue
		}
		a = r - c
		ans += a
		for i := 0; i < t && a > 0; i++ {
			if !vis[x-i] {
				vis[x-i] = true
				a--
			}
		}
	}
	fmt.Fprintln(out, ans)
}
