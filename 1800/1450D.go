package main

import (
	"bufio"
	"bytes"
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
	a := make([]int, n)
	cnt := make([]int, n)
	ans := bytes.Repeat([]byte{'0'}, n)
	ans[0] = '1'
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
		a[i]--
		cnt[a[i]]++
		if cnt[a[i]] > 1 {
			ans[0] = '0'
		}
	}

	l, r := 0, n-1
	for i, x := range cnt {
		if x == 0 {
			break
		}
		ans[n-i-1] = '1'
		if x > 1 || a[l] != i && a[r] != i {
			break
		}

		if a[l] == i {
			l++
		} else {
			r--
		}
	}

	fmt.Fprintln(out, string(ans))
}
