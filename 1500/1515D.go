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
	var n, l, r int
	fmt.Fscan(in, &n, &l, &r)
	a := make([]int, l)
	b := make([]int, r)
	cntA := make(map[int]int)
	cntB := make(map[int]int)
	same := 0
	for i := range a {
		fmt.Fscan(in, &a[i])
		cntA[a[i]]++
	}
	for i := range b {
		fmt.Fscan(in, &b[i])
		if cntA[b[i]] > 0 {
			cntA[b[i]]--
			same++
		} else {
			cntB[b[i]]++
		}
	}
	ans := 0
	if l > r {
		for _, v := range cntA {
			t := min(v/2, (l-r)/2)
			ans += t
			l -= t * 2
			if l == r {
				break
			}
		}
	} else if l < r {
		for _, v := range cntB {
			t := min(v/2, (r-l)/2)
			ans += t
			r -= t * 2
			if l == r {
				break
			}
		}
	}
	//fmt.Fprintln(out, ans)

	ans += max(l, r) - same

	fmt.Fprintln(out, ans)
}
