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

	var s []byte
	fmt.Fscan(in, &s)
	n := len(s)
	cnt, ans := 0, 0
	var a, b int
	h := &hp{}
	for i := 0; i < n; i++ {
		if s[i] == '(' {
			cnt++
		} else if s[i] == ')' {
			cnt--
		} else {
			cnt--
			fmt.Fscan(in, &a, &b)
			s[i] = ')'
			ans += b
			heap.Push(h, pair{b - a, i})
		}

		if cnt < 0 {
			if h.Len() == 0 {
				fmt.Fprintln(out, -1)
				return
			}
			cur := heap.Pop(h).(pair)
			ans -= cur.x
			s[cur.y] = '('
			cnt += 2
		}
	}

	if cnt != 0 {
		fmt.Fprintln(out, -1)
		return
	}
	fmt.Fprintln(out, ans)
	fmt.Fprintln(out, string(s))
}

type pair struct{ x, y int }
type hp []pair

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].x > h[j].x }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp) Pop() (v interface{}) {
	n := len(*h)
	pop := (*h)[n-1]
	*h = (*h)[:n-1]
	return pop
}
