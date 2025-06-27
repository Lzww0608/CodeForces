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
	nums := make([]int, n)
	for i := range nums {
		fmt.Fscan(in, &nums[i])
	}

	cnt, a, b := 0, 0, 0
	pre, round := 0, 0
	for l, r := 0, n-1; l <= r; round ^= 1 {
		cur := 0
		cnt++
		if round == 0 {
			for l <= r && cur <= pre {
				cur += nums[l]
				l++
			}
			a += cur
		} else {
			for l <= r && cur <= pre {
				cur += nums[r]
				r--
			}
			b += cur
		}
		pre = cur
	}

	fmt.Fprintln(out, cnt, a, b)
}
