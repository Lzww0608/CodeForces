package main

import (
	"bufio"
	. "fmt"
	"os"
	"sort"
)

func main() {

	var q int
	in := bufio.NewReader(os.Stdin)
	Fscan(in, &q)

	var n, l, r int
	for ; q > 0; q-- {
		Fscan(in, &n, &l, &r)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		sort.Ints(a)
		var ans int64
		for i, t := range a {
			x := sort.SearchInts(a[:i], r-t+1)
			y := sort.SearchInts(a[:i], l-t)
			ans += int64(x) - int64(y)
		}
		Println(ans)
	}
	return
}
