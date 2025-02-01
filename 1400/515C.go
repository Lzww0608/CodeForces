package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const N int = 10

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	factorial := make([]int, N)
	factorial[0] = 1
	for i := 1; i < N; i++ {
		factorial[i] = factorial[i-1] * i
	}

	var n int
	fmt.Fscan(in, &n)
	var s string
	fmt.Fscan(in, &s)

	ans := []int{}
	for i := 0; i < n; i++ {
		x := int(s[i] - '0')
		switch x {
		case 2:
			ans = append(ans, 2)
		case 3:
			ans = append(ans, 3)
		case 4:
			ans = append(ans, 3, 2, 2)
		case 5:
			ans = append(ans, 5)
		case 6:
			ans = append(ans, 5, 3)
		case 7:
			ans = append(ans, 7)
		case 8:
			ans = append(ans, 7, 2, 2, 2)
		case 9:
			ans = append(ans, 7, 3, 3, 2)
		}
	}
	sort.Slice(ans, func(i, j int) bool {
		return ans[i] > ans[j]
	})
	for _, v := range ans {
		fmt.Fprint(out, v)
	}
}
