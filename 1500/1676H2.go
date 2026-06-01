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

	var t int
	for fmt.Fscan(in, &t); t > 0; t-- {
		solve(in, out)
	}
}

func solve(in *bufio.Reader, out *bufio.Writer) {
	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}

	bit := NewBIT(n)
	ans := 0
	for i := n - 1; i >= 0; i-- {
		ans += bit.Sum(a[i])
		bit.Add(a[i], 1)
	}
	fmt.Fprintln(out, ans)
}

type BIT struct {
	n int
	f []int
}

func NewBIT(n int) *BIT {
	return &BIT{n: n, f: make([]int, n+1)}
}

func (b *BIT) Add(i, v int) {
	for i <= b.n {
		b.f[i] += v
		i += i & -i
	}
}

func (b *BIT) Sum(i int) int {
	s := 0
	for i > 0 {
		s += b.f[i]
		i -= i & -i
	}
	return s
}
