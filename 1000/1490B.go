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

	var t, n, x int
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &n)
		cnt := make([]int, 3)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &x)
			cnt[x%3]++
		}
		ans := 0

		for i := 0; min(cnt[0], cnt[1], cnt[2]) != n/3; i = (i + 1) % 3 {
			if cnt[i] > n/3 {
				ans += cnt[i] - n/3
				cnt[(i+1)%3] += cnt[i] - n/3
				cnt[i] = n / 3
			}
		}

		fmt.Fprintln(out, ans)
	}
}
