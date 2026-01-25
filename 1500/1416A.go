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
	pos := make([][]int, n+1)
	for i := range a {
		fmt.Fscan(in, &a[i])
		pos[a[i]] = append(pos[a[i]], i)
	}

	ans := make([]int, n)
	for i := range ans {
		ans[i] = -1
	}
	for i := 1; i <= n; i++ {
		if len(pos[i]) == 0 {
			continue
		}
		k := max(pos[i][0]+1, n-pos[i][len(pos[i])-1])
		for j := 1; j < len(pos[i]); j++ {
			k = max(k, pos[i][j]-pos[i][j-1])
		}

		for j := k - 1; j < n && ans[j] == -1; j++ {
			ans[j] = i
		}
	}

	for _, x := range ans {
		fmt.Fprintf(out, "%d ", x)
	}
	fmt.Fprintln(out)
}
