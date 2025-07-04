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
	var n, k, x int
	fmt.Fscan(in, &n, &k)
	cnt := map[int]int{}
	mx := 0
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &x)
		x = k - x%k
		if x != k {
			cnt[x]++
			if cnt[x] > cnt[mx] || cnt[x] == cnt[mx] && mx < x {
				mx = x
			}
		}
	}

	if mx == 0 {
		fmt.Fprintln(out, 0)
		return
	}

	fmt.Fprintln(out, k*(cnt[mx]-1)+mx+1)
}
