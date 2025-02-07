package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
	g := make([][]int, n)
	for i := 0; i < n-1; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		x--
		y--
		g[y] = append(g[y], x)
		g[x] = append(g[x], y)
	}
	var s string
	fmt.Fscan(in, &s)
	sum := strings.Count(s, "?")
	leaf, zero, one := 0, 0, 0
	for i := 1; i < n; i++ {
		if len(g[i]) == 1 {
			if s[i] == '?' {
				leaf++
			} else if s[i] == '0' {
				zero++
			} else {
				one++
			}
		}
	}

	if s[0] == '0' {
		fmt.Fprintln(out, one+(leaf+1)/2)
	} else if s[0] == '1' {
		fmt.Fprintln(out, zero+(leaf+1)/2)
	} else {
		if (sum-leaf)&1 == 0 && one == zero {
			fmt.Fprintln(out, one+(leaf+1)/2)
		} else {
			fmt.Fprintln(out, max(one, zero)+leaf/2)
		}
	}
}
