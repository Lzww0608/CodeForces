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

	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}

	cnt := make(map[int]bool)
	cnt[0] = true
	ans := 0
	sum := 0
	left := 0
	for l, r := 0, 0; l < n; l++ {
		for r < n {
			if cnt[sum+a[r]] {
				break
			}
			sum += a[r]
			r++

			cnt[sum] = true
		}

		ans += r - l
		cnt[left] = false
		left += a[l]
	}

	fmt.Fprintln(out, ans)
}
