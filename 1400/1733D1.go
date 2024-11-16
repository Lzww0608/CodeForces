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
	var n, x, y int
	fmt.Fscan(in, &n, &x, &y)
	var a, b string
	fmt.Fscan(in, &a, &b)
	cnt := []int{}
	for i := 0; i < n; i++ {
		if a[i] != b[i] {
			cnt = append(cnt, i)
		}
	}

	if len(cnt)&1 == 1 {
		fmt.Fprintln(out, -1)
		return
	}

	if len(cnt) != 2 || cnt[1] > cnt[0]+1 {
		fmt.Fprintln(out, y*len(cnt)/2)
	} else {
		fmt.Fprintln(out, min(x, y*2))
	}
}
