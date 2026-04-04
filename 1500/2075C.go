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
	a := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &a[i])
	}

	ans := 0
	sort.Ints(a)
	i, j := 0, m-1
	for k := 1; k < n; k++ {
		for i < m && a[i] < k {
			i++
		}
		for j >= 0 && a[j] >= n-k {
			j--
		}

		if i == m {
			break
		}

		t := min(m-i, m-j-1)
		ans += (m-j-1)*(m-i) - t
	}
	fmt.Fprintln(out, ans)
}
