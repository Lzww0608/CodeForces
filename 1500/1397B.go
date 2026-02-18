package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n)
	base := 0
	for i := range a {
		fmt.Fscan(in, &a[i])
		base += a[i] - 1
	}

	sort.Ints(a)
	ans := base
	for c := 1; ; c++ {
		sum, cur := 0, 1
		for _, x := range a {
			sum += abs(x - cur)
			if sum > base {
				cur = -1
				break
			}
			cur *= c
		}

		if cur == -1 {
			break
		}
		ans = min(ans, sum)
	}

	fmt.Fprintln(out, ans)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
