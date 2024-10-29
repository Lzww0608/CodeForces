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
	var n, m int
	fmt.Fscan(in, &n, &m)
	s := make([]string, n)
	for i := range s {
		fmt.Fscan(in, &s[i])
	}
	if s[0][0] == '1' {
		fmt.Fprintln(out, -1)
		return
	}
	a := make([][]int, n)
	ans := [][]int{}
	for i := range a {
		a[i] = make([]int, m)
		for j := range a[i] {
			a[i][j] = int(s[i][j] - '0')
		}

	}
	for j := m - 1; j >= 0; j-- {
		for i := n - 1; i >= 0; i-- {
			if a[i][j] == 1 {
				if i != 0 {
					ans = append(ans, []int{i, j + 1, i + 1, j + 1})
				} else {
					ans = append(ans, []int{i + 1, j, i + 1, j + 1})
				}
			}
		}

	}

	fmt.Fprintln(out, len(ans))
	for _, v := range ans {
		fmt.Fprintln(out, v[0], v[1], v[2], v[3])
	}
}
