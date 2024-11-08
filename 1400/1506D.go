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
	cnt := map[int]int{}
	mx := 0
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &x)
		cnt[x]++
		mx = max(mx, cnt[x])
	}
	if mx <= (n-1)/2 {
		if n&1 == 1 {
			fmt.Fprintln(out, 1)
		} else {
			fmt.Fprintln(out, 0)
		}
	} else {
		fmt.Fprintln(out, mx-(n-mx))
	}
}
