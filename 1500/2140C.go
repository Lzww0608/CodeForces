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
	for i := range a {
		fmt.Fscan(in, &a[i])
	}

	ans := 0
	for i, x := range a {
		if i&1 == 0 {
			ans += x
		} else {
			ans -= x
		}
	}

	odd, even := math.MinInt/3, math.MinInt/3
	tmp := ans
	for i, x := range a {
		ans = max(ans, tmp+i-i&1)
		if i&1 == 0 {
			ans = max(ans, tmp+i-2*x+odd)
			even = max(even, -i-x*2)
		} else {
			ans = max(ans, tmp+i+2*x+even)
			odd = max(odd, -i+x*2)
		}
	}

	fmt.Fprintln(out, ans)
}
