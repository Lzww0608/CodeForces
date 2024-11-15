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
	a := make([]int, n)
	id := 0
	for i := range a {
		fmt.Fscan(in, &a[i])
		if abs(a[id]) < abs(a[i]) {
			id = i
		}
	}
	if a[id] == 0 {
		fmt.Fprintln(out, 0)
		return
	}
	ans := [][]int{}
	if a[id] > 0 {
		ans = append(ans, []int{0, id})
		ans = append(ans, []int{0, id})
		for i := 0; i < n; i++ {
			ans = append(ans, []int{i, id})
		}
		for i := 1; i < n; i++ {
			ans = append(ans, []int{i, i - 1})
		}
	} else {
		ans = append(ans, []int{n - 1, id})
		ans = append(ans, []int{n - 1, id})
		for i := 0; i < n; i++ {
			ans = append(ans, []int{i, id})
		}
		for i := n - 2; i >= 0; i-- {
			ans = append(ans, []int{i, i + 1})
		}
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
