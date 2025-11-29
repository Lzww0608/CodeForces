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

	var nums []int

	var dfs func(int, int, int)
	dfs = func(a, b, x int) {
		if a > 5 || b > 5 {
			return
		}

		if a == b {
			nums = append(nums, x)
		}

		dfs(a+1, b, x*10+4)
		dfs(a, b+1, x*10+7)
	}
	dfs(0, 0, 0)

	sort.Ints(nums)

	var n int
	fmt.Fscan(in, &n)
	p := sort.SearchInts(nums, n)
	fmt.Fprintln(out, nums[p])
}
