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
	zero := 0
	neg := []int{}
	pos := []int{}
	for i := range a {
		fmt.Fscan(in, &a[i])
		if a[i] > 0 {
			pos = append(pos, a[i])
		} else if a[i] < 0 {
			neg = append(neg, a[i])
		} else {
			zero++
		}
	}
	zero = min(2, zero)

	if len(neg) > 2 || len(pos) > 2 {
		fmt.Fprintln(out, "NO")
		return
	} else if len(neg) == 0 && len(pos) == 0 {
		fmt.Fprintln(out, "YES")
		return
	}

	a = []int{}
	a = append(a, pos...)
	a = append(a, neg...)
	for i := 0; i < zero; i++ {
		a = append(a, 0)
	}
	vis := make(map[int]bool)
	for _, v := range a {
		vis[v] = true
	}

	m := len(a)
	for i := 0; i+2 < m; i++ {
		for j := i + 1; j < m-1; j++ {
			for k := j + 1; k < m; k++ {
				x := a[i] + a[j] + a[k]
				if !vis[x] {
					fmt.Fprintln(out, "NO")
					return
				}
			}
		}
	}

	fmt.Fprintln(out, "YES")
	return
}
