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
	a := make([][2]int, n)
	b1 := make([]int, n+1)
	b2 := make([]int, n+1)
	for i := range a {
		fmt.Fscan(in, &a[i][0], &a[i][1])
		if a[i][1] == 1 {
			b2[a[i][0]]++
		}
		b1[a[i][0]]++
	}

	cnt1 := make([]int, n+1)
	cnt2 := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		cnt1[b1[i]]++
		cnt2[b1[i]] = append(cnt2[b1[i]], b2[i])
	}

	ans1, ans2 := 0, 0
	x := 0
	h := &hp{}
	for i := n; i > 0; i-- {
		x += cnt1[i]
		for _, y := range cnt2[i] {
			heap.Push(h, y)
		}

		if x > 0 {
			ans1 += i
			x--
		}
		if h.Len() > 0 {
			y := heap.Pop(h).(int)
			ans2 += min(y, i)
		}
	}

	fmt.Fprintln(out, ans1, ans2)
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
	n := len(*h) - 1
	v = (*h)[n]
	*h = (*h)[:n]
	return v
}
