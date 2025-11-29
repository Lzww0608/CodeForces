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
	vis := make(map[int]bool)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	for _, x := range a {
		if x == 0 {
			continue
		}
		cnt, y := 0, 1
		for y*k <= x {
			y *= k
			cnt++
		}

		for x >= y && y > 0 {
			if vis[cnt] {
				fmt.Fprintln(out, "NO")
				return
			}
			vis[cnt] = true
			x -= y
			for x < y {
				y /= k
				cnt--
			}

		}
		if x != 0 {
			fmt.Fprintln(out, "NO")
			return
		}
	}

	fmt.Fprintln(out, "YES")
}
