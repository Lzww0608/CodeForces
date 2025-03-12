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
	var n, m, k int
	fmt.Fscan(in, &n, &m, &k)
	if n > m {
		n, m = m, n
	}
	if k < m {
		fmt.Fprintln(out, -1)
		return
	}

	ans := k
	if (k-n)&1 == 1 {
		ans--
	}
	if (k-m)&1 == 1 {
		ans--
	}
	fmt.Fprintln(out, ans)
}
