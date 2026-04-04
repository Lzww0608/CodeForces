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

	var t int
	for fmt.Fscan(in, &t); t > 0; t-- {
		solve(in, out)
	}
}

func solve(in *bufio.Reader, out *bufio.Writer) {
	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n)
	cnt := make(map[int]int)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
		cnt[a[i]]++
	}

	var q int
	for fmt.Fscan(in, &q); q > 0; q-- {
		var x, y int
		fmt.Fscan(in, &x, &y)

		D := x*x - 4*y
		if D < 0 {
			fmt.Fprint(out, "0 ")
			continue
		}

		r := int(math.Sqrt(float64(D)))
		for (r+1)*(r+1) <= D {
			r++
		}
		for r*r > D {
			r--
		}
		if r*r != D {
			fmt.Fprint(out, "0 ")
			continue
		}

		a := (x - r) / 2
		b := (x + r) / 2
		if r == 0 {
			fmt.Fprint(out, cnt[a]*(cnt[a]-1)/2, " ")
		} else {
			fmt.Fprint(out, cnt[a]*cnt[b], " ")
		}

	}
	fmt.Fprintln(out)
}
