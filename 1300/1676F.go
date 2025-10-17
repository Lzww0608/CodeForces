package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const MOD int = 1_000_000_007

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
	var n, k int
	fmt.Fscan(in, &n, &k)

	cnt := make(map[int]int)
	for range n {
		var x int
		fmt.Fscan(in, &x)
		cnt[x]++
	}

	a := make([]int, 0, len(cnt))
	for t := range cnt {
		a = append(a, t)
	}

	l, r := -1, -1
	d := 0
	sort.Ints(a)

	pre, cur := -1, 0
	for _, x := range a {
		if cnt[x] < k {
			pre, cur = -1, 0
			continue
		}

		if x != pre+1 {
			pre, cur = x, 1
		} else {
			pre, cur = x, cur+1
		}

		if cur > d {
			d = cur
			l, r = x-cur+1, x
		}
	}

	if d == 0 {
		fmt.Fprintln(out, -1)
	} else {
		fmt.Fprintln(out, l, r)
	}
}
