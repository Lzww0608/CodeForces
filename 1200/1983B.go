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
	var n, m int
	fmt.Fscan(in, &n, &m)
	a := make([]string, n)
	b := make([]string, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}
	for i := range b {
		fmt.Fscan(in, &b[i])
	}
	for i := 0; i < n; i++ {
		x, y := 0, 0
		for j := 0; j < m; j++ {
			x += int(a[i][j] - '0')
			y += int(b[i][j] - '0')
		}
		if x%3 != y%3 {
			fmt.Fprintln(out, "NO")
			return
		}

	}
	for j := 0; j < m; j++ {
		x, y := 0, 0
		for i := 0; i < n; i++ {
			x += int(a[i][j] - '0')
			y += int(b[i][j] - '0')
		}
		if x%3 != y%3 {
			fmt.Fprintln(out, "NO")
			return
		}
	}

	fmt.Fprintln(out, "YES")

}
