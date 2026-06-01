package main

import (
	"bufio"
	"fmt"
	"math"
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

var id = map[string]int{
	"BG": 0,
	"BR": 1,
	"BY": 2,
	"GR": 3,
	"GY": 4,
	"RY": 5,
}

var typ = []string{"BG", "BR", "BY", "GR", "GY", "RY"}

func solve(in *bufio.Reader, out *bufio.Writer) {
	var n, q int
	fmt.Fscan(in, &n, &q)
	a := make([]int, n)
	for i := range a {
		var s string
		fmt.Fscan(in, &s)
		if s[0] > s[1] {
			a[i] = id[string([]byte{s[1], s[0]})]
		} else {
			a[i] = id[s]
		}
	}

	pre := make([][6]int, n)
	nxt := make([][6]int, n)
	for i, v := range a {
		if i == 0 {
			pre[i] = [6]int{-1, -1, -1, -1, -1, -1}
		} else {
			pre[i] = pre[i-1]
		}

		pre[i][v] = i
	}

	for i := n - 1; i >= 0; i-- {
		if i == n-1 {
			nxt[i] = [6]int{-1, -1, -1, -1, -1, -1}
		} else {
			nxt[i] = nxt[i+1]
		}

		nxt[i][a[i]] = i
	}

	for range q {
		var i, j int
		fmt.Fscan(in, &i, &j)
		i--
		j--

		ans := math.MaxInt / 2
		si, sj := typ[a[i]], typ[a[j]]

		for x := range si {
			for y := range sj {
				if si[x] == sj[y] {
					ans = abs(i - j)
					break
				}

				for k, v := range id {
					if (k[0] == si[x] && k[1] == sj[y]) || (k[0] == sj[y] && k[1] == si[x]) {
						if pre[i][v] != -1 {
							ans = min(ans, abs(i-pre[i][v])+abs(j-pre[i][v]))
						}
						if nxt[i][v] != -1 {
							ans = min(ans, abs(i-nxt[i][v])+abs(j-nxt[i][v]))
						}
					}
				}
			}
		}

		if ans >= math.MaxInt/2 {
			ans = -1
		}
		fmt.Fprintln(out, ans)
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
