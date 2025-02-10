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

	var n, u, v int
	fmt.Fscan(in, &n)
	f := make([]int, n*2)
	l := make([]int, n*2)
	r := make([]int, n*2)
	for i := range f {
		f[i] = i
	}
	var find func(int) int
	find = func(x int) int {
		if f[x] != x {
			f[x] = find(f[x])
		}
		return f[x]
	}

	for i := 1; i < n; i++ {
		fmt.Fscan(in, &u, &v)
		fu, fv := find(u), find(v)
		l[n+i] = fu
		r[n+i] = fv
		f[fu], f[fv] = n+i, n+i
	}

	var output func(int)
	output = func(x int) {
		if l[x] == 0 {
			fmt.Fprint(out, x, " ")
		} else {
			output(l[x])
			output(r[x])
		}
	}

	output(n*2 - 1)
}
