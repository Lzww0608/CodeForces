package main

import (
	"bufio"
	"fmt"
	"os"
)

const N int = 200_005

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, k, q int
	fmt.Fscan(in, &n, &k, &q)
	dif := make([]int, N+1)

	var l, r int
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &l, &r)
		dif[l]++
		dif[r+1]--
	}

	for i := 1; i < N; i++ {
		dif[i] += dif[i-1]
	}

	pre := make([]int, N)
	for i := 1; i < N; i++ {
		if dif[i] >= k {
			pre[i] = pre[i-1] + 1
		} else {
			pre[i] = pre[i-1]
		}
	}

	for i := 0; i < q; i++ {
		fmt.Fscan(in, &l, &r)
		fmt.Fprintln(out, pre[r]-pre[l-1])
	}

}
