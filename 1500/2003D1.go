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
	a := make([][]int, n)

	k := 0
	for i := range a {
		var l int
		fmt.Fscan(in, &l)
		a[i] = make([]int, l)
		cnt := make([]bool, l+2)
		for j := range a[i] {
			fmt.Fscan(in, &a[i][j])
			if a[i][j] <= l {
				cnt[a[i][j]] = true
			}
		}

		t := 0
		for t <= l && cnt[t] {
			t++
		}
		t++
		for t <= l && cnt[t] {
			t++
		}

		k = max(k, t)
	}

	if m <= k {
		fmt.Fprintln(out, k*(m+1))
	} else {
		fmt.Fprintln(out, k*(k+1)+(m-k)*(k+1+m)/2)
	}
}
