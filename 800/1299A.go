package main

import (
	"bufio"
	. "fmt"
	"math"
	"os"
)

func main() {

	var n int
	in := bufio.NewReader(os.Stdin)
	Fscan(in, &n)

	a := make([]int, n)
	pre := make([]int, n+1)
	pre[0] = math.MaxInt32
	for i := range a {
		Fscan(in, &a[i])
		pre[i+1] = pre[i] & ^a[i]
	}

	suf := math.MaxInt32
	res, idx := -1, 0
	for i := n - 1; i >= 0; i-- {
		t := a[i] & pre[i] & suf
		if t >= res {
			res, idx = t, i
		}
		suf &= ^a[i]
	}
	Print(a[idx], " ")
	for i, x := range a {
		if i != idx {
			Print(x, " ")
		}
	}
	return
}
