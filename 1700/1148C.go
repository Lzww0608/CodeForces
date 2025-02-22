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
	cnt := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
		a[i]--
		cnt[a[i]] = i
	}

	ans := [][2]int{}
	for i := 0; i < n; i++ {
		if a[i] == i {
			continue
		}
		j := cnt[i]
		if abs(i-j) >= n/2 {
			ans = append(ans, [2]int{i, j})
		} else if i >= n/2 {
			ans = append(ans, [2]int{0, i})
			ans = append(ans, [2]int{0, j})
			ans = append(ans, [2]int{0, i})
		} else if n-j-1 >= n/2 {
			ans = append(ans, [2]int{i, n - 1})
			ans = append(ans, [2]int{j, n - 1})
			ans = append(ans, [2]int{i, n - 1})
		} else {
			ans = append(ans, [2]int{i, n - 1})
			ans = append(ans, [2]int{0, n - 1})
			ans = append(ans, [2]int{0, j})
			ans = append(ans, [2]int{0, n - 1})
			ans = append(ans, [2]int{i, n - 1})
		}
		cnt[i], cnt[a[i]] = i, j
		a[i], a[j] = a[j], a[i]
	}

	fmt.Fprintln(out, len(ans))
	for _, v := range ans {
		fmt.Fprintln(out, v[0]+1, v[1]+1)
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
