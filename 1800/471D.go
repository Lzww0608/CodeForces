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

	var n, w int
	fmt.Fscan(in, &n, &w)
	a := make([]int, n+1)
	c := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
		c[i] = a[i] - a[i-1]
	}
	b := make([]int, w+1)
	d := make([]int, w+1)
	for i := 1; i <= w; i++ {
		fmt.Fscan(in, &b[i])
		d[i] = b[i] - b[i-1]
	}

	fmt.Fprintln(out, kmp(c[2:], d[2:]))
	return
}

func kmp(text, pattern []int) (ans int) {
	m := len(pattern)
	if m == 0 {
		return len(text) + 1
	}
	pi := make([]int, m)
	cnt := 0
	for i := 1; i < m; i++ {
		v := pattern[i]
		for cnt > 0 && pattern[cnt] != v {
			cnt = pi[cnt-1]
		}
		if pattern[cnt] == v {
			cnt++
		}
		pi[i] = cnt
	}

	cnt = 0
	for i := 0; i < len(text); i++ {
		for cnt > 0 && pattern[cnt] != text[i] {
			cnt = pi[cnt-1]
		}
		if pattern[cnt] == text[i] {
			cnt++
		}
		if cnt == m {
			ans++
			cnt = pi[cnt-1]
		}
	}

	return
}
