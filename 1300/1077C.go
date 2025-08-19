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
	vis := make(map[int]int)
	sum := 0
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
		sum += a[i]
		vis[a[i]]++
	}

	ans := make([]int, 0, n)
	for i, x := range a {
		cur := sum - x
		vis[x]--
		if cur&1 == 0 && vis[cur/2] > 0 {
			ans = append(ans, i+1)
		}
		vis[x]++
	}

	fmt.Fprintln(out, len(ans))
	for _, x := range ans {
		fmt.Fprintf(out, "%d ", x)
	}
}
