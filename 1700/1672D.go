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
	b := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &b[i])
	}

	i, j := 0, 0
	cnt := make(map[int]int)
	for j < n {
		if i < n && j < n && a[i] == b[j] {
			i++
			j++
			continue
		}

		if cnt[b[j]] > 0 && b[j] == b[j-1] {
			cnt[b[j]]--
			j++
		} else if i < n {
			cnt[a[i]]++
			i++
		} else {
			fmt.Fprintln(out, "NO")
			return
		}
	}
	fmt.Fprintln(out, "YES")
}
