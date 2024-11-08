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
	a := make([]int, n)
	cnt := map[int]struct{}{}
	for i := range a {
		fmt.Fscan(in, &a[i])
		cnt[a[i]] = struct{}{}
	}

	if len(cnt) > k {
		fmt.Fprintln(out, -1)
		return
	}
	b := []int{}
	for i := range cnt {
		b = append(b, i)
	}
	fmt.Fprintln(out, n*k)
	for i := 0; i < n; i++ {
		for _, x := range b {
			fmt.Fprint(out, x, " ")
		}
		for j := 0; j < k-len(b); j++ {
			fmt.Fprint(out, b[0], " ")
		}
	}
	fmt.Fprintln(out)
}
