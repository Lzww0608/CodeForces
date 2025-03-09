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

	var n, m, k, x, y int
	fmt.Fscan(in, &n, &m, &k, &x, &y)
	if n == 1 || n == 2 {
		t := k / (n * m)
		mx, mn, cur := t+1, t, t
		if k%(n*m) == 0 {
			mx--
		}
		if k%(n*m) >= (x-1)*m+y {
			cur++
		}
		fmt.Fprintln(out, mx, mn, cur)
		return
	}
	t := k / (m*n + (n-2)*m)
	mx, mn, cur := t*2, t, t
	q := k % (m*n + (n-2)*m)
	if q > 0 {
		if t == 0 {
			mx++
		} else if q > m {
			mx++
		}

		if q > m*n {
			mx++
		}
		if q >= m*n {
			mn++
		}
	}

	if x == 1 {
		if q >= y {
			cur++
		}
	} else if x == n {
		if q >= (x-1)*m+y {
			cur++
		}
	} else {
		cur *= 2
		if q >= (x-1)*m+y {
			cur++
		}
		if q >= m*n+(n-x-1)*m+y {
			cur++
		}
	}

	fmt.Fprintln(out, mx, mn, cur)
}
