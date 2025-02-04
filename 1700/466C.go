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
	ans, sum := 0, 0
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
		sum += a[i]
	}
	if sum%3 != 0 {
		fmt.Fprintln(out, 0)
		return
	}
	sum /= 3
	cnt, cur := 0, 0
	for i, x := range a[:n-1] {
		cur += x

		if i > 0 && cur == sum*2 {
			ans += cnt
		}

		if cur == sum {
			cnt++
		}
	}
	fmt.Fprintln(out, ans)
}
