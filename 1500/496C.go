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
	s := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &s[i])
	}
	p := make(map[int]bool)
	for i := 1; i < n; i++ {
		p[i] = true
	}

	ans := 0
	for j := 0; j < m; j++ {
		f := false
		for i := 1; i < n; i++ {
			if !p[i] {
				continue
			}
			if s[i][j] < s[i-1][j] {
				ans++
				f = false
				break
			} else if s[i][j] > s[i-1][j] {
				f = true
			}
		}

		if f {
			for i := 1; i < n; i++ {
				if s[i][j] > s[i-1][j] {
					p[i] = false
				}
			}
		}
	}

	fmt.Fprintln(out, ans)
}
