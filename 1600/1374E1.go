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

	var n, k int
	fmt.Fscan(in, &n, &k)
	a, b, c := hp{}, hp{}, hp{}
	for range n {
		var t, x, y int
		fmt.Fscan(in, &t, &x, &y)
		if x == 1 {
			if y == 1 {
				heap.Push(&a, t)
			} else {
				heap.Push(&b, t)
			}
		} else if y == 1 {
			heap.Push(&c, t)
		}
	}

	ans := 0
	for range k {
		if a.Len() == 0 {
			if b.Len() == 0 || c.Len() == 0 {
				fmt.Fprintln(out, -1)
				return
			}
			x := heap.Pop(&b).(int)
			y := heap.Pop(&c).(int)
			ans += x + y
		} else {
			if b.Len() == 0 || c.Len() == 0 {
				ans += heap.Pop(&a).(int)
			} else {
				x, y := b[0], c[0]
				if a[0] <= x+y {
					ans += heap.Pop(&a).(int)
				} else {
					ans += heap.Pop(&c).(int) + heap.Pop(&b).(int)
				}
			}
		}
	}

	fmt.Fprintln(out, ans)
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
func (h *hp) Pop() interface{} {
	a := *h
	v := a[len(a)-1]
	*h = a[:len(a)-1]
	return v
}
