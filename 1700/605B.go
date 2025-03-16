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

	var n, m int
	fmt.Fscan(in, &n, &m)
	a := make([]int, 0, m)
	b := make([]int, 0, m)
	s := make([]int, m)
	for i := 0; i < m; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		s[i] = u
		if v == 1 {
			a = append(a, i)
		} else {
			b = append(b, i)
		}
	}
	sort.Slice(a, func(i, j int) bool { return s[a[i]] < s[a[j]] })
	sort.Slice(b, func(i, j int) bool { return s[b[i]] < s[b[j]] })

	edges := make([][2]int, m)
	pos := make(map[int]int)
	for j, i := range a {
		edges[i] = [2]int{n, j + 1}
		pos[j+1] = s[i]
	}

	x, y := 1, 2
	for _, i := range b {
		for y <= n-1 && s[i] < pos[y] {
			x, y = y, y+1
		}
		if y >= n {
			fmt.Fprintln(out, -1)
			return
		}
		edges[i] = [2]int{x, y}
		if x == 1 {
			y++
			x = y - 1
		} else {
			x--
		}
	}

	for _, v := range edges {
		fmt.Fprintln(out, v[0], v[1])
	}
	return
}
