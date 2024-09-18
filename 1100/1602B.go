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

	var t, n, q, x, k int
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			fmt.Fscan(in, &a[i])
		}
		b := [][]int{}
		b = append(b, a)
		for f := true; f; {
			cnt := map[int]int{}
			for _, x := range b[len(b)-1] {
				cnt[x]++
			}
			tmp := make([]int, n)
			eq := 0
			for i := 0; i < n; i++ {
				tmp[i] = cnt[b[len(b)-1][i]]
				if tmp[i] == b[len(b)-1][i] {
					eq++
				}
			}
			b = append(b, tmp)
			f = eq != n
		}

		fmt.Fscan(in, &q)
		for i := 0; i < q; i++ {
			fmt.Fscan(in, &x, &k)
			k = min(k, len(b)-1)
			fmt.Fprintln(out, b[k][x-1])
		}
	}
}
