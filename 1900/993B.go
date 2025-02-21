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

	var n, m int
	fmt.Fscan(in, &n, &m)
	possible := make([][2]map[int]int, max(m, n))
	for i := 0; i < max(m, n); i++ {
		possible[i][0] = make(map[int]int)
		possible[i][1] = make(map[int]int)
	}
	allpossible := make(map[int]int)

	a := make([][2]int, n)
	b := make([][2]int, m)

	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i][0], &a[i][1])
	}

	for i := 0; i < m; i++ {
		fmt.Fscan(in, &b[i][0], &b[i][1])
	}

	solve := func(a, b [2]int) int {
		cnt, val := 0, -1
		for i := 0; i < 2; i++ {
			for j := 0; j < 2; j++ {
				if a[i] == b[j] {
					cnt++
					val = a[i]
				}
			}
		}
		if cnt == 1 {
			return val
		}
		return -1
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			val := solve(a[i], b[j])
			if val != -1 {
				allpossible[val]++
				possible[j][1][val]++
				possible[i][0][val]++
				if len(possible[j][1]) > 1 || len(possible[i][0]) > 1 {
					fmt.Fprintln(out, -1)
					return
				}
			}
		}
	}

	if len(allpossible) == 1 {
		for k := range allpossible {
			fmt.Fprintln(out, k)
			return
		}
	}

	fmt.Fprintln(out, 0)
	return
}
