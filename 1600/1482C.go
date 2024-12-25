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
	fmt.Fscan(in, &n, &m)

	a := make([][]int, m)
	cnt := make([]int, n+1)
	ans := make([]int, m)
	for i := range a {
		fmt.Fscan(in, &k)
		a[i] = make([]int, k)
		for j := range a[i] {
			fmt.Fscan(in, &a[i][j])
		}
		cnt[a[i][0]]++
		ans[i] = a[i][0]
	}

	f := 0
	for i := 1; i <= n; i++ {
		if cnt[i] > (m+1)/2 {
			f = i
			break
		}
	}

	if f != 0 {
		for i := range a {
			if a[i][0] != f || len(a[i]) == 1 {
				continue
			} else if cnt[f] <= (m+1)/2 {
				break
			}
			cnt[f]--
			ans[i] = a[i][1]
			cnt[a[i][1]]++
		}
	}

	if cnt[f] > (m+1)/2 {
		fmt.Fprintln(out, "NO")
		return
	}

	fmt.Fprintln(out, "YES")
	for _, x := range ans {
		fmt.Fprint(out, x, " ")
	}
	fmt.Fprintln(out)
}
