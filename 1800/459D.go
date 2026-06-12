package main

import (
	"bufio"
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
	cnt := make(map[int]int)
	left := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
		cnt[a[i]]++
		left[i] = cnt[a[i]]
	}

	b := NewBIT(n)
	ans := 0
	clear(cnt)
	for i := n - 1; i >= 0; i-- {
		cnt[a[i]]++
		ans += b.Sum(left[i] - 1)
		b.Add(cnt[a[i]], 1)
	}

	fmt.Fprintln(out, ans)
}

type BIT struct {
	n    int
	tree []int
}

func NewBIT(n int) *BIT {
	return &BIT{n: n, tree: make([]int, n+1)}
}

func (b *BIT) Add(i, v int) {
	for i <= b.n {
		b.tree[i] += v
		i += i & -i
	}
}

func (b *BIT) Sum(i int) int {
	s := 0
	for i > 0 {
		s += b.tree[i]
		i -= i & -i
	}
	return s
}
