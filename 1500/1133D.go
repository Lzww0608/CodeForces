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

	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}
	b := make([]int, n)
	for i := range b {
		fmt.Fscan(in, &b[i])
	}

	add := 0
	type pair struct {
		x, y int
	}
	cnt := make(map[pair]int)
	for i := range n {
		x, y := abs(a[i]), abs(b[i])
		if x == 0 {
			if y == 0 {
				add++
			}
			continue
		}
		g := gcd(x, y)
		x /= g
		y /= g
		if a[i]*b[i] < 0 {
			cnt[pair{-x, y}]++
		} else {
			cnt[pair{x, y}]++
		}
	}

	ans := 0
	for _, v := range cnt {
		ans = max(ans, v)
	}
	fmt.Fprintln(out, ans+add)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
