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
	var n, m int
	fmt.Fscan(in, &n, &m)
	cnt := make([][2]int, n+m-1)
	a := make([][]int, n)
	for i := range a {
		a[i] = make([]int, m)
		for j := range a[i] {
			fmt.Fscan(in, &a[i][j])
			d := i + j
			if a[i][j] == 0 {
				cnt[d][0]++
			} else {
				cnt[d][1]++
			}
		}
	}

	ans := 0
	for i := 0; i < (n+m-1)/2; i++ {
		j := n + m - 2 - i
		ans += min(cnt[i][0]+cnt[j][0], cnt[i][1]+cnt[j][1])
	}

	fmt.Fprintln(out, ans)
}
