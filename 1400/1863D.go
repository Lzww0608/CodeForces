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
	var n, m int
	fmt.Fscan(in, &n, &m)
	str := make([]string, n)
	ans := make([][]byte, n)
	for i := 0; i < n; i++ {
		ans[i] = make([]byte, m)
		for j := range ans[i] {
			ans[i][j] = '.'
		}
		fmt.Fscan(in, &str[i])
	}

	for i := 0; i < n; i++ {
		b, w := 0, 0
		for j := 0; j < m; j++ {
			if str[i][j] == '.' {
				continue
			}
			if str[i][j] == 'U' {
				if w >= b {
					ans[i][j] = 'B'
					b++
				} else {
					ans[i][j] = 'W'
					w++
				}
			}
		}
		for j := 0; j < m; j++ {
			if str[i][j] == 'D' {
				if ans[i-1][j] == 'B' {
					ans[i][j] = 'W'
					w++
				} else {
					ans[i][j] = 'B'
					b++
				}
			}
		}
		if b != w {
			fmt.Fprintln(out, -1)
			return
		}
	}

	for j := 0; j < m; j++ {
		b, w := 0, 0
		for i := 0; i < n; i++ {
			if str[i][j] == 'L' {
				if w >= b {
					ans[i][j] = 'B'
					b++
				} else {
					ans[i][j] = 'W'
					w++
				}
			}
		}
		for i := 0; i < n; i++ {
			if str[i][j] == 'R' {
				if ans[i][j-1] == 'B' {
					ans[i][j] = 'W'
					w++
				} else {
					ans[i][j] = 'B'
					b++
				}
			}
		}
		if b != w {
			fmt.Fprintln(out, -1)
			return
		}
	}

	for i := 0; i < n; i++ {
		fmt.Fprintln(out, string(ans[i]))
	}
}
