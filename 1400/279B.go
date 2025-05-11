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

	var n, m int
	fmt.Fscan(in, &n, &m)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	ans := 0
	sum := 0
	for l, r := 0, 0; r < n; r++ {
		sum += a[r]
		for sum > m {
			sum -= a[l]
			l++
		}

		ans = max(ans, r-l+1)
	}

	fmt.Fprintln(out, ans)
}
