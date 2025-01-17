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

	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n)
	sum := 0
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
		sum += a[i]
	}
	if sum&1 == 1 {
		fmt.Fprintln(out, 0)
		return
	}
	//fmt.Fprintln(out, sum, sum/2)

	f := make([]bool, sum/2+1)
	f[0] = true
	for _, x := range a {
		for i := len(f) - 1; i >= x; i-- {
			if f[i-x] {
				f[i] = true
			}
		}
	}

	if !f[len(f)-1] {
		fmt.Fprintln(out, 0)
		return
	}

	ans, val := -1, math.MaxInt
	for i, x := range a {
		t := x & (-x)
		if t < val {
			val = t
			ans = i + 1
		}
	}
	fmt.Fprintln(out, 1)
	fmt.Fprintln(out, ans)
}
