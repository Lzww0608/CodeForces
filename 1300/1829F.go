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
	var n, m int
	fmt.Fscan(in, &n, &m)

	cnt := make([]int, n+1)
	for i := 0; i < m; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		cnt[u]++
		cnt[v]++
	}

	cnts := make(map[int]int)
	for i := 1; i <= n; i++ {
		if cnt[i] != 0 {
			cnts[cnt[i]]++
		}
	}

	a := []int{}
	for _, v := range cnts {
		a = append(a, v)
	}
	sort.Ints(a)

	if len(a) == 3 {
		fmt.Fprintln(out, a[1], a[2]/a[1])
	} else {
		fmt.Fprintln(out, a[0]-1, a[1]/(a[0]-1))
	}
}
