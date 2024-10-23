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
	ans, cnt := n, 1
	for i := 2; i*i <= n; i++ {
		t, k := i*i, 1
		for n%t == 0 {
			t *= i
			k++
		}
		if k > cnt {
			cnt = k
			ans = i
		}
	}

	fmt.Fprintln(out, cnt)
	for i := 0; i < cnt-1; i++ {
		fmt.Fprint(out, ans, " ")
		n /= ans
	}
	fmt.Fprintln(out, n)
}
