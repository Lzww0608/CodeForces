package main

import (
	"bufio"
	"fmt"
	"os"
)

const N int = 1001

var F [N][N]int

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
	var n, q int
	fmt.Fscan(in, &n, &q)

	for i := range N {
		for j := range N {
			F[i][j] = 0
		}
	}

	var h, w int
	for range n {
		fmt.Fscan(in, &h, &w)
		F[h][w] += h * w
	}

	for i := 1; i < N; i++ {
		for j := 1; j < N; j++ {
			F[i][j] += F[i-1][j] + F[i][j-1] - F[i-1][j-1]
		}
	}

	for range q {
		var h1, w1, h2, w2 int
		fmt.Fscan(in, &h1, &w1, &h2, &w2)
		ans := F[h2-1][w2-1] - F[h1][w2-1] - F[h2-1][w1] + F[h1][w1]
		fmt.Fprintln(out, ans)
	}

}
