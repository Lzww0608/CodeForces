package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

type Item struct {
	r, l int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].r < pq[j].r || pq[i].r == pq[j].r && pq[i].l < pq[j].l
}

func (pq *PriorityQueue) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Item))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var t int
	fmt.Fscan(in, &t)

	for t > 0 {
		t--
		var n int
		fmt.Fscan(in, &n)

		pq := make(PriorityQueue, 0)
		heap.Push(&pq, &Item{-n, 0})

		cnt := 1
		ans := make([]int, n)

		for pq.Len() > 0 {
			item := heap.Pop(&pq).(*Item)
			r, l := item.r, item.l
			r = l - r - 1

			mid := (l + r) / 2
			ans[mid] = cnt
			cnt++

			if mid > l {
				heap.Push(&pq, &Item{l - mid, l})
			}
			if mid < r {
				heap.Push(&pq, &Item{mid - r, mid + 1})
			}
		}

		for i := 0; i < n; i++ {
			fmt.Fprintf(out, "%d ", ans[i])
		}
		fmt.Fprintln(out)
	}
}
