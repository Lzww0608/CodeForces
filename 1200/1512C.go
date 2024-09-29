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

	var t, a, b int
next:
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &a, &b)
		var s string
		fmt.Fscan(in, &s)
		ans := []byte(s)
		n := a + b
		x, y := 0, 0
		for l, r := 0, n-1; l < r; l, r = l+1, r-1 {
			if ans[l] != ans[r] {
				if ans[l] != '?' && ans[r] != '?' {
					fmt.Fprintln(out, -1)
					continue next
				} else if ans[l] != '?' {
					ans[r] = ans[l]
				} else {
					ans[l] = ans[r]
				}

				if ans[l] == '0' {
					x += 2
				} else {
					y += 2
				}
			} else {
				if ans[l] == '0' {
					x += 2
				} else if ans[l] == '1' {
					y += 2
				}
			}

		}

		for l, r := 0, n-1; l < r; l, r = l+1, r-1 {
			if ans[l] == ans[r] {
				if ans[l] == '?' {
					if x+1 < a {
						ans[l], ans[r] = '0', '0'
						x += 2
					} else {
						ans[l], ans[r] = '1', '1'
						y += 2
					}
				}
			}
		}

		if n&1 == 1 {
			if s[n/2] == '0' {
				x++
			} else if s[n/2] == '1' {
				y++
			} else {
				if x < a {
					x++
					ans[n/2] = '0'
				} else {
					y++
					ans[n/2] = '1'
				}
			}
		}

		if x > a || y > b {
			fmt.Fprintln(out, -1)
			continue next
		}

		fmt.Fprintln(out, string(ans))
	}
}
