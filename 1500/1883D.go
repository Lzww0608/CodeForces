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

	var q int
	fmt.Fscan(in, &q)

	a := make(map[int]int)
	b := make(map[int]int)
	left := hp{}
	right := hp{}
	for i := 0; i < q; i++ {
		var s string
		var ll, rr int
		fmt.Fscan(in, &s, &ll, &rr)
		if s == "+" {
			a[ll]++
			b[rr]++
			heap.Push(&left, ll)
			heap.Push(&right, -rr)
		} else {
			a[ll]--
			b[rr]--
		}

		for left.Len() > 0 && a[left[0]] == 0 {
			heap.Pop(&left)
		}
		for right.Len() > 0 && b[-right[0]] == 0 {
			heap.Pop(&right)
		}

		if left.Len() > 0 && right.Len() > 0 && left[0] > -right[0] {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
}

type hp []int

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i] > h[j] }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(x any) {
	*h = append(*h, x.(int))
}
func (h *hp) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
