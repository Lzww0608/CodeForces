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

	var n, m int
	fmt.Fscan(in, &n, &m)
	var s, t string
	fmt.Fscan(in, &s)
	fmt.Fscan(in, &t)

	left := make([]int, m)
	right := make([]int, m)

	l := 0
	for i := range n {
		if s[i] == t[l] {
			left[l] = i
			l++
			if l == m {
				break
			}
		}
	}

	r := m - 1
	for i := n - 1; i >= 0; i-- {
		if s[i] == t[r] {
			right[r] = i
			r--
			if r == 0 {
				break
			}
		}
	}

	ans := 1
	for i := range m - 1 {
		ans = max(ans, right[i+1]-left[i])
	}

	fmt.Fprintln(out, ans)
}
