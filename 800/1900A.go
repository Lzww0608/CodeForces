package main

import "fmt"

func main() {
	var t, n int
next:
	for fmt.Scan(&t); t > 0; t-- {
		fmt.Scan(&n)
		var s string
		fmt.Scan(&s)
		s += "#"
		ans, cnt := 0, 0
		for i := range s {
			if s[i] == '#' {
				if cnt >= 3 {
					fmt.Println(2)
					continue next
				} else {
					ans += cnt
					cnt = 0
				}
			} else {
				cnt++
			}
		}
		fmt.Println(ans)
	}
}
