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
	var s string
	fmt.Fscan(in, &s)
	x, y := 0, 0
	cnt := make([]int, 4)
	for i := range s {
		if s[i] == 'N' {
			y++
			cnt[0]++
		} else if s[i] == 'S' {
			y--
			cnt[1]++
		} else if s[i] == 'E' {
			x++
			cnt[2]++
		} else {
			x--
			cnt[3]++
		}
	}

	if x&1 == 1 || y&1 == 1 || n&1 == 1 || (n == 2 && (x == 0 && y == 0)) {
		fmt.Fprintln(out, "NO")
		return
	}
	ans := make([]byte, n)
	x = x / 2
	y = y / 2
	sum := 0
	for i := range s {
		if x > 0 && s[i] == 'E' || x < 0 && s[i] == 'W' {
			ans[i] = 'R'
			sum++
			if x < 0 {
				x++
			} else {
				x--
			}
		} else if y > 0 && s[i] == 'N' || y < 0 && s[i] == 'S' {
			ans[i] = 'R'
			sum++
			if y < 0 {
				y++
			} else {
				y--
			}
		}
	}
	if sum == 0 {
		if cnt[2] > 0 {
			for i := range s {
				if s[i] == 'E' && sum <= 0 {
					ans[i] = 'R'
					sum++
					if sum == 0 {
						break
					}
				} else if s[i] == 'W' && sum >= 0 {
					ans[i] = 'R'
					sum--
					if sum == 0 {
						break
					}
				}
			}
		} else {
			for i := range s {
				if s[i] == 'S' && sum <= 0 {
					ans[i] = 'R'
					sum++
					if sum == 0 {
						break
					}
				} else if s[i] == 'N' && sum >= 0 {
					ans[i] = 'R'
					sum--
					if sum == 0 {
						break
					}
				}
			}
		}

	}
	for i := range ans {
		if ans[i] != 'R' {
			ans[i] = 'H'
		}
	}
	fmt.Fprintln(out, string(ans))
}
