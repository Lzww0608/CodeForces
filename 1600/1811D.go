package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const N int = 50

var F [N]int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	F[0], F[1] = 1, 1
	for i := 2; i < N; i++ {
		F[i] = F[i-1] + F[i-2]
	}

	var t int
	for fmt.Fscan(in, &t); t > 0; t-- {
		solve(in, out)
	}
}

func solve(in io.Reader, out io.Writer) {
	var n, x, y int
	fmt.Fscan(in, &n, &x, &y)

	for n > 0 {
		if y <= F[n] && y > F[n-1] {
			fmt.Fprintln(out, "NO")
			return
		} else if y > F[n] {
			y -= F[n]
		}
		x, y = y, x
		n -= 1
	}
	fmt.Fprintln(out, "YES")
}
