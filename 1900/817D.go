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
	a := make([]int, n+1)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	a[n] = math.MaxInt32

	ans := 0
	st := []int{-1}
	for i, x := range a {
		for len(st) > 1 && x > a[st[len(st)-1]] {
			k := st[len(st)-1]
			st = st[:len(st)-1]
			ans += a[k] * (k - st[len(st)-1]) * (i - k)
		}
		st = append(st, i)
	}

	st = []int{-1}
	a[n] = math.MinInt32
	for i, x := range a {
		for len(st) > 1 && x < a[st[len(st)-1]] {
			k := st[len(st)-1]
			st = st[:len(st)-1]
			ans -= a[k] * (k - st[len(st)-1]) * (i - k)
		}
		st = append(st, i)
	}

	fmt.Fprintln(out, ans)
}
