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

	var t int
	for fmt.Fscan(in, &t); t > 0; t-- {
		solve(in, out)
	}
}

func solve(in *bufio.Reader, out *bufio.Writer) {
	var n, m, d int
	fmt.Fscan(in, &n, &m, &d)
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}
	cost, ans, sum := 0, 0, 0
	h := &hp{}
	for _, x := range a {
		cost -= d
		if x > 0 {
			sum += x
			heap.Push(h, x)
			if h.Len() > m {
				sum -= heap.Pop(h).(int)
			}
			ans = max(ans, sum+cost)
		}
	}

	fmt.Fprintln(out, ans)
	return
}

type hp []int

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i] < h[j] }
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
	return
}
