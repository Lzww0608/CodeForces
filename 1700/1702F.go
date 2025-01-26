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

	// 处理 a 数组，保存奇数形式的计数
	cnt := make(map[int]int)
	for i := 0; i < n; i++ {
		var x int
		fmt.Fscan(in, &x)
		for x%2 == 0 {
			x >>= 1
		}
		cnt[x]++
	}

	// 使用最大堆模拟 C++ 的 multiset 行为
	h := &IntHeap{}
	heap.Init(h)
	for i := 0; i < n; i++ {
		var x int
		fmt.Fscan(in, &x)
		heap.Push(h, x)
	}

	for h.Len() > 0 {
		x := heap.Pop(h).(int)
		if cnt[x] > 0 {
			cnt[x]--
			if cnt[x] == 0 {
				delete(cnt, x)
			}
		} else {
			if x == 1 {
				// 无法再分割，直接失败
				fmt.Fprintln(out, "NO")
				return
			}
			// 分割后重新插入堆
			heap.Push(h, x/2)
		}
	}

	fmt.Fprintln(out, "YES")
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] } // 最大堆
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
