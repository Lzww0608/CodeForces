package main

import (
	. "fmt"
)

func main() {

	var t, n, m, r, c int
next:
	for Scanf("%d\n", &t); t > 0; t-- {
		Scanf("%d %d %d %d\n", &n, &m, &r, &c)
		r--
		c--
		a := make([]string, n)
		for i := 0; i < n; i++ {
			Scanf("%s\n", &a[i])
		}
		if a[r][c] == 'B' {
			Println(0)
			continue
		}
		for i := 0; i < m; i++ {
			if a[r][i] == 'B' {
				Println(1)
				continue next
			}
		}
		for i := 0; i < n; i++ {
			if a[i][c] == 'B' {
				Println(1)
				continue next
			}
		}
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				if a[i][j] == 'B' {
					Println(2)
					continue next
				}
			}
		}
		Println(-1)
	}

}
