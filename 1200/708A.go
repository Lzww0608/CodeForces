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

	var s string
	fmt.Fscan(in, &s)
	n := len(s)
	ans := []byte(s)
	l := 0
	for l < n && ans[l] == 'a' {
		l++
	}
	r := l
	for r < n && ans[r] != 'a' {
		r++
	}
	if l == n {
		ans[n-1] = 'z'
	}
	for i := l; i < r; i++ {
		ans[i] -= 1
	}

	fmt.Fprintln(out, string(ans))
}
