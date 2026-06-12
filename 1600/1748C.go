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
	for i := range a {
		fmt.Fscan(in, &a[i])
	}

	ans, mx := 0, 0
	f := false
	sum := 0
	cnt := make(map[int]int)
	for _, x := range a {
		if x == 0 {
			if !f {
				f = true
				ans += cnt[0]
			} else {
				ans += mx
			}

			mx = 0
			clear(cnt)
		}

		sum += x
		cnt[sum]++
		if cnt[sum] > mx {
			mx = cnt[sum]
		}
	}

	if f {
		ans += mx
	} else {
		ans += cnt[0]
	}

	fmt.Fprintln(out, ans)
}
