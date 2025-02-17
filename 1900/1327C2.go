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

	var n int
	fmt.Fscan(in, &n)
	a := make([]node, n)
	for i := range a {
		fmt.Fscan(in, &a[i].x, &a[i].y, &a[i].z)
		a[i].id = i + 1
	}
	sort.Slice(a, func(i, j int) bool {
		if a[i].x == a[j].x {
			if a[i].y == a[j].y {
				return a[i].z < a[j].z
			}
			return a[i].y < a[j].y
		}
		return a[i].x < a[j].x
	})
	vis := make([]bool, n+1)
	l := 0
	for r := 1; r < n; r++ {
		if l != -1 && a[r].x == a[l].x && a[r].y == a[l].y {
			vis[r] = true
			vis[l] = true
			fmt.Fprintln(out, a[l].id, a[r].id)
			l = -1
		} else {
			l = r
		}
	}

	l = 0
	for vis[l] {
		l++
	}
	for r := l + 1; r < n; r++ {
		if vis[r] {
			continue
		}
		if l != -1 && a[r].x == a[l].x {
			vis[r] = true
			vis[l] = true
			fmt.Fprintln(out, a[l].id, a[r].id)
			l = -1
		} else {
			l = r
		}
	}

	l = 0
	for vis[l] {
		l++
	}
	for r := l + 1; r < n; r++ {
		if vis[r] {
			continue
		}
		if l != -1 {
			vis[r] = true
			vis[l] = true
			fmt.Fprintln(out, a[l].id, a[r].id)
			l = -1
		} else {
			l = r
		}
	}
}

type node struct {
	x, y, z, id int
}
