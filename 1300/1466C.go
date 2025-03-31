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
	var s string
	fmt.Fscan(in, &s)
	ans := 0
	n := len(s)
	vis := make(map[int]bool)
	for i := 0; i < n-1; i++ {
		if vis[i] {
			continue
		}
		if i+2 < n && s[i] == s[i+2] {
			ans++
			vis[i+2] = true
		}
		if s[i] == s[i+1] && !vis[i+1] {
			ans++
			vis[i+1] = true
		}
	}

	fmt.Fprintln(out, ans)
}
