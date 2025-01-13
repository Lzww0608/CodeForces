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
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	ans := [][]int{}
	for i := 0; i < n; i++ {
		for j := 0; j < n-1; j++ {
			if a[j] < a[j+1] {
				ans = append(ans, []int{a[j], a[j+1]})
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n-1; j++ {
			if a[j] > a[j+1] {
				ans = append(ans, []int{a[j], a[j+1]})
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	fmt.Fprintln(out, len(ans))
	for i := len(ans) - 1; i >= 0; i-- {
		fmt.Fprintln(out, ans[i][0], ans[i][1])
	}
}
