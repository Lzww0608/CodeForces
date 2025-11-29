package main

import (
	"bufio"
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
	cnt := make(map[int]int)
	for i := range a {
		fmt.Fscan(in, &a[i])
		cnt[a[i]]++
	}

	m := make(map[int]int)
	for _, v := range cnt {
		m[v]++
	}

	b := make([]int, 0, len(m))
	for k := range m {
		b = append(b, k)
	}
	sort.Ints(b)

	ans, sum, pre := n, n, 0
	cur := 0
	for _, x := range b {
		sum -= x * m[x]
		cur += m[x]
		ans = min(ans, pre+sum-x*(len(cnt)-cur))
		pre += x * m[x]
	}

	fmt.Fprintln(out, ans)
}
