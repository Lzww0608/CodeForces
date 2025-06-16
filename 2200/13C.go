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

	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	var ans int
	h := &hp{}
	for _, x := range a {
		if h.Len() > 0 && (*h)[0] > x {
			cur := heap.Pop(h).(int)
			ans += cur - x
			heap.Push(h, x)
		}
		heap.Push(h, x)
	}

	fmt.Fprintln(out, ans)
}

type hp []int

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i] > h[j] }
func (h hp) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *hp) Push(v interface{}) {
	*h = append(*h, v.(int))
}
func (h *hp) Pop() (v interface{}) {
	n := len(*h)
	v = (*h)[n-1]
	*h = (*h)[:n-1]
	return v
}
