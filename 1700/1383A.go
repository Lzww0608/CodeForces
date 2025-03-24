package main

import (
	"bufio"
	"fmt"
	"os"
)

const N int = 20

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
	var s, t string
	fmt.Fscan(in, &s, &t)
	a := [N][N]int{}

	for i := 0; i < n; i++ {
		if s[i] > t[i] {
			fmt.Fprintln(out, -1)
			return
		}
		x, y := int(s[i]-'a'), int(t[i]-'a')
		if x == y {
			continue
		}
		a[x][y]++
	}

	ans := 0
	for j := 1; j < N; j++ {
		for i := 0; i < j; i++ {
			if a[i][j] > 0 {
				ans++
				for k := j + 1; k < N; k++ {
					a[j][k] += a[i][k]
					a[i][k] = 0
				}
			}
		}
	}
	fmt.Fprintln(out, ans)
}
