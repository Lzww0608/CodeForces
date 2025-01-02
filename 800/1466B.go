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
	a := make([]int, n)
	m := make(map[int]int)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
		m[a[i]]++
	}

	ans := len(m)
	for i := 0; i < n; i++ {
		if m[a[i]] > 1 {
			m[a[i]+1]++
			if m[a[i]+1] == 1 {
				ans++
			}
		}
	}
	fmt.Fprintln(out, ans)
}
