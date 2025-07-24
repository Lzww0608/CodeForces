package main

import (
	"bufio"
	"fmt"
	"os"
)

const N int = 31

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
	var n, k int
	fmt.Fscan(in, &n, &k)

	cnt := [N]int{}
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
		for j := 0; j < N; j++ {
			if a[i]&(1<<j) == 0 {
				cnt[j]++
			}
		}
	}

	ans := 0
	for i := N - 1; i >= 0; i-- {
		if k >= cnt[i] {
			k -= cnt[i]
			ans |= 1 << i
		}
	}
	fmt.Fprintln(out, ans)
}
