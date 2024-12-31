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
	k--
	if n <= 50 && k >= (1<<(n-1)) {
		fmt.Fprintln(out, -1)
		return
	}

	p := make([]int, n)
	l, r := 0, n-1
	for i := 1; i < n; i++ {
		if n-i-1 > 50 || k < (1<<(n-i-1)) {
			p[l] = i
			l++
		} else {
			p[r] = i
			r--
			k -= 1 << (n - i - 1)
		}
	}
	p[l] = n
	for i := 0; i < n; i++ {
		fmt.Fprint(out, p[i], " ")
	}
	fmt.Fprintln(out)
}
