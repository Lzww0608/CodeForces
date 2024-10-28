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
	fmt.Fscan(in, &n)
	a := make([][]int, n)
	cnt := map[int]int{}
	for i := range a {
		fmt.Fscan(in, &m)
		a[i] = make([]int, m)
		for j := range a[i] {
			fmt.Fscan(in, &a[i][j])
			cnt[a[i][j]]++
		}
	}
next:
	for _, v := range a {
		for _, x := range v {
			if cnt[x] == 1 {
				continue next
			}
		}
		fmt.Fprintln(out, "Yes")
		return
	}
	fmt.Fprintln(out, "No")
}
