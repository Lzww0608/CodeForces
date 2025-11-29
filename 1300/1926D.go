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
	var n, x int
	fmt.Fscan(in, &n)
	cnt := make(map[int]int)
	ans := n

	mask := (1 << 31) - 1
	for range n {
		fmt.Fscan(in, &x)
		y := mask ^ x
		if cnt[y] > 0 {
			cnt[y]--
			ans--
		} else {
			cnt[x]++
		}
	}

	fmt.Fprintln(out, ans)
}
