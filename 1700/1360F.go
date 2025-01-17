package main

import (
	"bufio"
	"fmt"
	"os"
)

const N int = 26

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
	var n, m int
	fmt.Fscan(in, &n, &m)
	s := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &s[i])
	}
	ans := []byte(s[0])

	for i := 0; i < m; i++ {
		c := ans[i]
	next:
		for j := 0; j < N; j++ {
			ans[i] = byte('a' + j)
			for t := 0; t < n; t++ {
				cnt := 0
				for k := 0; k < m; k++ {
					if ans[k] != s[t][k] {
						if cnt++; cnt > 1 {
							break
						}
					}
				}
				if cnt > 1 {
					continue next
				}
			}
			fmt.Fprintln(out, string(ans))
			return
		}
		ans[i] = c
	}
	fmt.Fprintln(out, -1)
}
