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
	a := make([]string, n)
	cnt := 0
	for i := range a {
		fmt.Fscan(in, &a[i])
		for j := range a[i] {
			cnt += int(a[i][j] - '0')
		}
	}
	mx := 0
	for j := 0; j < n; j++ {
		cur := 0
		for i := 0; i < n; i++ {
			if a[i][(i+j)%n] == '1' {
				cur++
			}
		}
		mx = max(mx, cur)
	}

	fmt.Fprintln(out, cnt-mx+n-mx)
}
