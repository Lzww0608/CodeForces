package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, l, r int
	fmt.Fscan(in, &n, &l, &r)
	a := make([]int, n)
	p := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &p[i])
	}

	mx, mn := math.MinInt, math.MaxInt
	b := make([]int, n)
	for i := 0; i < n; i++ {
		b[i] = a[i] + p[i]
		if b[i] > mx {
			mx = b[i]
		}
		if b[i] < mn {
			mn = b[i]
		}
	}
	if mx-mn > r-l {
		fmt.Fprintln(out, -1)
		return
	}

	d := 0
	if l > mn {
		d = l - mn
	} else {
		d = r - mx
	}

	for i := 0; i < n; i++ {
		fmt.Fprint(out, b[i]+d, " ")
	}
	fmt.Fprintln(out)
}
