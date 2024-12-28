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
	s := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &s[i])
		if s[i] != "removeMin" {
			fmt.Fscan(in, &a[i])
		}
	}

	h := &hp{}
	op := []string{}
	b := []int{}
	for i := 0; i < n; i++ {
		if s[i] == "insert" {
			op = append(op, "insert")
			b = append(b, a[i])
			heap.Push(h, a[i])
		} else if s[i] == "removeMin" {
			if h.Len() == 0 {
				op = append(op, "insert")
				b = append(b, int(-1e9))
				heap.Push(h, int(-1e9))
			}
			op = append(op, "removeMin")
			b = append(b, 0)
			heap.Pop(h)
		} else {
			for h.Len() > 0 && (*h)[0] < a[i] {
				t := heap.Pop(h).(int)
				op = append(op, "removeMin")
				b = append(b, t)
			}
			if h.Len() == 0 || (*h)[0] != a[i] {
				op = append(op, "insert")
				b = append(b, a[i])
				heap.Push(h, a[i])
			}
			op = append(op, "getMin")
			b = append(b, a[i])
		}
	}

	fmt.Fprintln(out, len(op))
	for i := range op {
		if op[i] != "removeMin" {
			fmt.Fprintln(out, op[i], b[i])
		} else {
			fmt.Fprintln(out, op[i])
		}
	}
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
	item := (*h)[n-1]
	*h = (*h)[:n-1]
	return item
}
