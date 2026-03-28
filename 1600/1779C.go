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
	var n, m int
	fmt.Fscan(in, &n, &m)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(in, &a[i])
	}

	ans := calc(a[m:])
	b := make([]int, 0, m)
	for i := m - 1; i > 0; i-- {
		b = append(b, -a[i])
	}
	ans += calc(b)
	fmt.Fprintln(out, ans)
}

func calc(a []int) (res int) {
	h := &hp{}
	sum := 0
	for _, x := range a {
		sum += x
		heap.Push(h, x)
		for sum < 0 {
			sum -= 2 * heap.Pop(h).(int)
			res++
		}
	}

	return
}

type hp []int

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i] < h[j] }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *hp) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
