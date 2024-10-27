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
	cnt := 0
	for i := range a {
		fmt.Fscan(in, &a[i])
		if a[i] != 0 {
			cnt++
		}
	}
	if cnt&1 == 1 {
		fmt.Fprintln(out, -1)
		return
	}
	ans := [][]int{}
	for i := 0; i < n; i++ {
		if a[i] == 0 {
			j := i + 1
			for j < n && a[j] == 0 {
				j++
			}
			ans = append(ans, []int{i + 1, j})
			i = j - 1
		} else {
			j := i + 1
			for j < n && a[j] == 0 {
				j++
			}
			t := j - i + 1
			if a[i] == a[j] {
				if t&1 == 0 {
					ans = append(ans, []int{i + 1, j + 1})
				} else {
					ans = append(ans, []int{i + 1, i + 1})
					ans = append(ans, []int{i + 2, j + 1})
				}
			} else {
				if t&1 == 1 {
					ans = append(ans, []int{i + 1, j + 1})
				} else {
					ans = append(ans, []int{i + 1, j})
					ans = append(ans, []int{j + 1, j + 1})
				}
			}
			i = j
		}
	}

	fmt.Fprintln(out, len(ans))
	for _, v := range ans {
		fmt.Fprintln(out, v[0], v[1])
	}
}
