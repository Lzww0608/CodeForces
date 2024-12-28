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
	var n, q, d int
	fmt.Fscan(in, &n, &q)
	for i := 1; i < n; i++ {
		fmt.Fprintln(out, i, i+1)
	}
	a := make([]int, n)
	b := []int{1}
	for i := range a {
		a[i] = i
	}
	for i := 0; i < q; i++ {
		fmt.Fscan(in, &d)
		d++
		if len(a) == d || len(b)+1 == d {
			fmt.Fprintln(out, -1, -1, -1)
			continue
		}

		if len(a) < d {
			tmp := make([]int, d-len(a))
			for j := range tmp {
				tmp[j] = b[len(b)-len(tmp)+j]
			}
			b = b[:len(b)-len(tmp)]
			fmt.Fprintln(out, tmp[0]+1, b[len(b)-1]+1, a[len(a)-1]+1)
			for j := range tmp {
				a = append(a, tmp[j])
			}
		} else {
			tmp := make([]int, len(a)-d)
			for j := range tmp {
				tmp[j] = a[len(a)-len(tmp)+j]
			}
			a = a[:len(a)-len(tmp)]
			fmt.Fprintln(out, tmp[0]+1, a[len(a)-1]+1, b[len(b)-1]+1)
			for j := range tmp {
				b = append(b, tmp[j])
			}
		}
	}
}
