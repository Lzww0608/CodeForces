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

const N int = 26

func solve(in *bufio.Reader, out *bufio.Writer) {
	var n int
	fmt.Fscan(in, &n)
	var s string
	fmt.Fscan(in, &s)

	if n&1 == 0 {
		ans, mx := 0, 0
		cnt := make(map[byte]int)
		for i := 0; i < n; i += 2 {
			cnt[s[i]]++
			if cnt[s[i]] > mx {
				mx = cnt[s[i]]
			}
		}
		ans = n/2 - mx
		mx = 0
		clear(cnt)
		for i := 1; i < n; i += 2 {
			cnt[s[i]]++
			if cnt[s[i]] > mx {
				mx = cnt[s[i]]
			}
		}
		ans += n/2 - mx
		fmt.Fprintln(out, ans)
		return
	}

	ans := n
	p1, p2 := [N]int{}, [N]int{}
	cnt1 := make([][N]int, n)
	cnt2 := make([][N]int, n)
	for i := range s {
		x := int(s[i] - 'a')
		if i&1 == 0 {
			p1[x]++
			cnt1[i] = p1
		} else {
			p2[x]++
			cnt2[i] = p2
		}
	}

	rev1 := make([][N]int, n+2)
	rev2 := make([][N]int, n+2)
	p1, p2 = [N]int{}, [N]int{}
	for i := n - 1; i >= 0; i-- {
		x := int(s[i] - 'a')
		if i&1 == 0 {
			p1[x]++
			rev1[i] = p1
		} else {
			p2[x]++
			rev2[i] = p2
		}
	}

	for i := range n {
		a, b := 0, 0
		if i-1 >= 0 {
			for j := range N {
				x, y := 0, 0
				if (i-1)&1 == 0 {
					x = cnt1[i-1][j]
				} else {
					x = cnt2[i-1][j]
				}
				if (i+2)&1 == 0 {
					y = rev1[i+2][j]
				} else {
					y = rev2[i+2][j]
				}
				a = max(a, x+y)
			}
		} else {
			for j := range N {
				if (i+2)&1 == 0 {
					a = max(a, rev1[i+2][j])
				} else {
					a = max(a, rev2[i+2][j])
				}
			}
		}

		if i-2 >= 0 {
			for j := range N {
				x, y := 0, 0
				if (i-2)&1 == 0 {
					x = cnt1[i-2][j]
				} else {
					x = cnt2[i-2][j]
				}
				if (i+1)&1 == 0 {
					y = rev1[i+1][j]
				} else {
					y = rev2[i+1][j]
				}
				b = max(b, x+y)
			}
		} else {
			for j := range N {
				if (i+1)&1 == 0 {
					b = max(b, rev1[i+1][j])
				} else {
					b = max(b, rev2[i+1][j])
				}
			}
		}

		//fmt.Fprintln(out, i, a, b)
		ans = min(ans, n-a-b)
	}
	fmt.Fprintln(out, ans)
}
