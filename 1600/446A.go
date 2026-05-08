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

	left := make([]int, n)
	right := make([]int, n)

	for i := range a {
		left[i] = 1
		if i > 0 && a[i] > a[i-1] {
			left[i] = left[i-1] + 1
		}
	}
	for i := n - 1; i >= 0; i-- {
		right[i] = 1
		if i+1 < n && a[i] < a[i+1] {
			right[i] = right[i+1] + 1
		}
	}

	ans := 1
	for i := range a {
		if left[i] > ans {
			ans = left[i]
		}
		if i > 0 && left[i-1]+1 > ans {
			ans = left[i-1] + 1
		}
		if i+1 < n && right[i+1]+1 > ans {
			ans = right[i+1] + 1
		}
		if i > 0 && i+1 < n && a[i+1] > a[i-1]+1 {
			cur := left[i-1] + 1 + right[i+1]
			if cur > ans {
				ans = cur
			}
		}
	}
	if ans > n {
		ans = n
	}

	fmt.Fprintln(out, ans)
}
