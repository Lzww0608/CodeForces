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
	a := make([]int, n)
	pre := make([]int, n+1)
	suf := make([]int, n+1)
	for i := range a {
		fmt.Fscan(in, &a[i])
		suf[a[i]]++
	}
	mex1, mex2 := 0, 0
	for suf[mex1] > 0 {
		mex1++
	}

	for i := 0; i < n; i++ {
		pre[a[i]]++
		suf[a[i]]--
		for pre[mex2] > 0 {
			mex2++
		}
		if suf[a[i]] == 0 && mex1 > a[i] {
			mex1 = a[i]
		}

		if mex1 == mex2 {
			fmt.Fprintln(out, 2)
			fmt.Fprintln(out, 1, i+1)
			fmt.Fprintln(out, i+2, n)
			return
		}
	}
	fmt.Fprintln(out, -1)

}
