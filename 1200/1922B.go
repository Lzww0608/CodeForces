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

	var t, n int
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &n)
		a := make([]int, n)
		cnt := map[int]int{}
		for i := range a {
			fmt.Fscan(in, &a[i])
			cnt[a[i]]++
		}
		ans := 0
		sum := 0
		for i := 0; i <= n; i++ {
			ans += cnt[i] * (cnt[i] - 1) * (cnt[i] - 2) / 6
			ans += sum * cnt[i] * (cnt[i] - 1) / 2
			sum += cnt[i]
		}
		fmt.Fprintln(out, ans)
	}
}
