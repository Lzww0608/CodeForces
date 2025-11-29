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

	solve(in, out)

}

func solve(in *bufio.Reader, out *bufio.Writer) {
	var n, k int
	fmt.Fscan(in, &n, &k)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	m := make(map[int]bool)
	st := []int{}
	for _, x := range a {
		if m[x] {
			continue
		}

		if len(st) == k {
			m[st[0]] = false
			st = st[1:]
		}
		st = append(st, x)
		m[x] = true
	}

	fmt.Fprintln(out, len(st))
	for i := len(st) - 1; i >= 0; i-- {
		fmt.Fprintf(out, "%d ", st[i])
	}
}
