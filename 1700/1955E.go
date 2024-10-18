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
	var s string
	fmt.Fscan(in, &s)
	a := make([]int, n)
	for i := range a {
		if s[i] == '1' {
			a[i] = 1
		}
	}
next:
	for k := n; k > 1; k-- {
		cnt := 0
		b := make([]int, n)
		for i := 0; i < n; i++ {
			if i >= k && b[i-k] == 1 {
				cnt--
			}
			if (cnt&1)+a[i] != 1 {
				if i > n-k {
					continue next
				}
				b[i] = 1
				cnt++
			}
		}
		fmt.Fprintln(out, k)
		return
	}
	fmt.Fprintln(out, 1)
}
