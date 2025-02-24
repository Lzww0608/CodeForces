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
	var n, m, k int
	fmt.Fscan(in, &n, &m, &k)
	s := make([]string, n)
	ans := make([][]byte, n)
	cnt := 0
	for i := 0; i < n; i++ {
		ans[i] = make([]byte, m)
		fmt.Fscan(in, &s[i])
		for j := range s[i] {
			if s[i][j] == 'R' {
				cnt++
			}
		}
	}
	avg, rem := cnt/k, cnt%k
	t, sum := byte('0'), 0
	for i := 0; i < n; i++ {
		if i&1 == 0 {
			for j := 0; j < m; j++ {
				if s[i][j] == 'R' {
					sum++
					if rem > 0 && sum == avg+2 || rem == 0 && sum == avg+1 {
						sum = 1
						if rem > 0 {
							rem--
						}
						if t == '9' {
							t = 'a'
						} else if t == 'z' {
							t = 'A'
						} else {
							t++
						}
					}
				}

				ans[i][j] = t
			}
		} else {
			for j := m - 1; j >= 0; j-- {
				if s[i][j] == 'R' {
					sum++
					if rem > 0 && sum == avg+2 || rem == 0 && sum == avg+1 {
						sum = 1
						if rem > 0 {
							rem--
						}
						if t == '9' {
							t = 'a'
						} else if t == 'z' {
							t = 'A'
						} else {
							t++
						}
					}
				}

				ans[i][j] = t
			}
		}
	}

	for i := 0; i < n; i++ {
		fmt.Fprintln(out, string(ans[i]))
	}
}
