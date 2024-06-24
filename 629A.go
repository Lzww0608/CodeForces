package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	s := make([]string, n)
	ans := 0
	for i := range s {
		fmt.Scan(&s[i])
	}
	cnt := make([]int, n)
	for i := range s {
		x := 0
		for j := range s {
			if s[i][j] == 'C' {
				x++
				cnt[j]++
			}
		}
		ans += x * (x - 1) / 2
	}
	for _, x := range cnt {
		ans += x * (x - 1) / 2
	}
	fmt.Println(ans)
}
