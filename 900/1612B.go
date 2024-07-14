package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t, n, a, b int
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &n, &a, &b)
		nums := []int{a}
		for i := n; i >= 1; i-- {
			if i != a && i != b {
				nums = append(nums, i)
			}
		}
		nums = append(nums, b)
		x := slices.Min(nums[:n/2])
		y := slices.Max(nums[n/2:])
		if x == a && y == b {
			for _, x := range nums {
				fmt.Fprint(out, x, " ")
			}
			fmt.Fprintln(out)
		} else {
			fmt.Fprintln(out, -1)
		}
	}
}
