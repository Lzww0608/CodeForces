package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
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
	b := make([]int, n)

	one := &hp{}
	two := &hp{}
	for i := range a {
		fmt.Fscan(in, &a[i])
		b[i] = a[i]
		if a[i] == 1 {
			heap.Push(one, i)
		} else if a[i] == 2 {
			heap.Push(two, i)
		}
	}

	sort.Ints(b)
	ans := [][2]int{}
	for i := n - 1; i >= 0; i-- {
		if a[i] == b[i] {
			continue
		}

		if a[i] == 1 && b[i] == 2 {
			j := heap.Pop(two).(int)
			ans = append(ans, [2]int{j + 1, i + 1})
			heap.Push(one, j)
			a[j] = 1
		} else if a[i] == 0 && b[i] == 1 {
			j := heap.Pop(one).(int)
			ans = append(ans, [2]int{j + 1, i + 1})
			a[j] = 0
		} else {
			j := heap.Pop(two).(int)
			k := heap.Pop(one).(int)
			ans = append(ans, [2]int{k + 1, i + 1})
			ans = append(ans, [2]int{j + 1, i + 1})
			a[j] = 1
			a[k] = 0
			heap.Push(one, j)
		}
	}

	fmt.Fprintln(out, len(ans))
	for _, p := range ans {
		fmt.Fprintln(out, p[0], p[1])
	}
}

type hp []int

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i] < h[j] }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(int)) }
func (h *hp) Pop() interface{}   { v := (*h)[len(*h)-1]; *h = (*h)[:len(*h)-1]; return v }
