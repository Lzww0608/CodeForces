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

	var t int
	for fmt.Fscan(in, &t); t > 0; t-- {
		var s string
		fmt.Fscan(in, &s)
		n := len(s)
		x, y := int(s[0]-'a'), int(s[n-1]-'a')
		x, y = min(x, y), max(x, y)
		cost := abs(x - y)
		cnt := 2
		a := [][]int{}
		for i, c := range s[1 : n-1] {
			cur := int(c - 'a')
			if cur >= x && cur <= y {
				cnt++
				a = append(a, []int{cur, i + 2})
			}
		}
		sort.Slice(a, func(i, j int) bool {
			return a[i][0] < a[j][0]
		})

		fmt.Fprintln(out, cost, cnt)
		fmt.Fprint(out, 1, " ")
		if int(s[0]-'a') <= int(s[n-1]-'a') {
			for _, v := range a {
				fmt.Fprint(out, v[1], " ")
			}
		} else {
			m := len(a)
			for i := m - 1; i >= 0; i-- {
				fmt.Fprint(out, a[i][1], " ")
			}
		}
		fmt.Fprintln(out, n)
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
