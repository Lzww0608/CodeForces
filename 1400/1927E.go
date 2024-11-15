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
	ans := make([]int, n)
	l, r := 1, n
	for i := 0; i < k; i++ {
		for j := i; j < n; j += k {
			if j&1 == 0 {
				ans[j] = l
				l++
			} else {
				ans[j] = r
				r--
			}
		}
	}
	for _, x := range ans {
		fmt.Fprint(out, x, " ")
	}
	fmt.Fprintln(out)
}
