package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, k1, k2 int
	fmt.Fscan(in, &n, &k1, &k2)
	k := k1 + k2
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}

	h := hp{}
	cnt := make(map[int]int)
	for i := range n {
		var x int
		fmt.Fscan(in, &x)
		cnt[abs(a[i]-x)]++
	}

	for k, v := range cnt {
		if k == 0 {
			continue
		}
		heap.Push(&h, pair{k, v})
	}

	for k > 0 && h.Len() > 0 {
		cur := heap.Pop(&h).(pair)
		x, c := cur.x, cur.c
		if h.Len() == 0 {
			t := k / c
			if t < x {
				rem := k % c
				k = 0
				heap.Push(&h, pair{x - t - 1, rem})
				heap.Push(&h, pair{x - t, c - rem})
			} else {
				k -= x * c
			}
			break
		} else {
			next := h[0]
			y := next.x
			if k >= (x-y)*c {
				k -= (x - y) * c
				h[0].c += c
			} else {
				t := k / c
				rem := k % c
				heap.Push(&h, pair{x - t, c - rem})
				if rem != 0 {
					heap.Push(&h, pair{x - t - 1, rem})
				}
				k = 0
				break
			}
		}
	}

	if k > 0 {
		fmt.Fprintln(out, k&1)
		return
	}

	ans := 0
	for _, v := range h {
		ans += v.x * v.x * v.c
	}

	fmt.Fprintln(out, ans)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

type pair struct {
	x, c int
}
type hp []pair

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].x > h[j].x }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(x any) {
	*h = append(*h, x.(pair))
}
func (h *hp) Pop() (x any) {
	old := *h
	n := len(old)
	x = old[n-1]
	*h = old[:n-1]
	return
}
