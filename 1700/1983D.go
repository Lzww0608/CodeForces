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
	b := make([]int, n)
	c := make([]int, n)
	d := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
		c[i] = a[i]
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &b[i])
		d[i] = b[i]
	}
	sort.Ints(c)
	sort.Ints(d)
	for i := 0; i < n; i++ {
		if c[i] != d[i] {
			fmt.Fprintln(out, "NO")
			return
		}
	}

	pos := make(map[int]int)
	for i := 0; i < n; i++ {
		pos[b[i]] = i
	}

	cnt := 0
	for i := 0; i < n; i++ {
		if a[i] != b[i] {
			j := pos[a[i]]
			pos[b[i]], pos[a[i]] = pos[a[i]], pos[b[i]]
			b[i], b[j] = b[j], b[i]
			cnt++
		}
	}

	if cnt&1 == 1 {
		fmt.Fprintln(out, "NO")
	} else {
		fmt.Fprintln(out, "YES")
	}
}
