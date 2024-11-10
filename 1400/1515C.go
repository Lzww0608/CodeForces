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
	var n, m, x int
	fmt.Fscan(in, &n, &m, &x)
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}
	h := &hp{}
	for i := 1; i <= m; i++ {
		heap.Push(h, []int{0, i})
	}

	id := 0
	ans := make([]int, n)
	for id < n {
		t := heap.Pop(h).([]int)
		ans[id] = t[1]
		t[0] += a[id]
		heap.Push(h, t)
		id++
	}
	fmt.Fprintln(out, "YES")
	for _, x := range ans {
		fmt.Fprint(out, x, " ")
	}
	fmt.Fprintln(out)
}

type hp [][]int

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i][0] < h[j][0] }
func (h hp) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *hp) Push(v interface{}) {
	*h = append(*h, v.([]int))
}
func (h *hp) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
