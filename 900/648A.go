package main

import (
	"fmt"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	s := make([]string, n)
	a := make([]int, m)
	for i := range s {
		fmt.Scan(&s[i])
	}
	for j := 0; j < m; j++ {
		cnt := 0
		for i := n - 1; i >= 0; i-- {
			if s[i][j] == '*' {
				cnt++
			} else {
				break
			}
		}
		a[j] = cnt
	}
	up, down := 0, 0
	for i := 1; i < m; i++ {
		if a[i] > a[i-1] {
			up = max(up, a[i]-a[i-1])
		} else if a[i] < a[i-1] {
			down = max(down, a[i-1]-a[i])
		}
	}
	fmt.Println(up, down)
}
