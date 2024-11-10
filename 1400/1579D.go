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
	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n)
	h := &hp{}
	for i := range a {
		fmt.Fscan(in, &a[i])
		if a[i] > 0 {
			heap.Push(h, []int{a[i], i + 1})
		}
	}
	ans := [][]int{}
	for h.Len() > 1 {
		v1 := heap.Pop(h).([]int)
		v2 := heap.Pop(h).([]int)
		ans = append(ans, []int{v1[1], v2[1]})
		if v1[0]--; v1[0] > 0 {
			heap.Push(h, v1)
		}
		if v2[0]--; v2[0] > 0 {
			heap.Push(h, v2)
		}
	}
	fmt.Fprintln(out, len(ans))
	for _, v := range ans {
		fmt.Fprintln(out, v[0], v[1])
	}
}

type hp [][]int

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i][0] > h[j][0] }
func (h hp) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *hp) Push(v interface{}) {
	*h = append(*h, v.([]int))
}
func (h *hp) Pop() (v interface{}) {
	n := len(*h)
	v = (*h)[n-1]
	*h = (*h)[:n-1]
	return
}
