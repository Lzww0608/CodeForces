package main

import (
	"bufio"
	"fmt"
	"os"
)

const N int = 1 << 8

var a [N][N]int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	cur := 0
	for i := 0; i < N; i += 2 {
		for j := 0; j < N; j += 2 {
			a[i][j], a[i][j+1] = cur, cur+1
			a[i+1][j], a[i+1][j+1] = cur+2, cur+3
			cur += 4
		}
	}

	var t int
	for fmt.Fscan(in, &t); t > 0; t-- {
		solve(in, out)
	}
}

func solve(in *bufio.Reader, out *bufio.Writer) {
	var n, m int
	fmt.Fscan(in, &n, &m)
	fmt.Fprintln(out, n*m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Fprint(out, a[i][j], " ")
		}
		fmt.Fprintln(out)
	}

}
