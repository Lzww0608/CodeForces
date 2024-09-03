package main

import (
	"fmt"
	"strings"
)

func main() {
	var m, n int
	fmt.Scan(&m, &n)
	a := make([]string, m)
	last := 0
	for i := range a {
		fmt.Scan(&a[i])
		if strings.Contains(a[i], "W") {
			last = i
		}
	}

	cur := 0
	ans := 0
	for i := 0; i <= last; i++ {
		if !strings.Contains(a[i], "W") {
			continue
		}
		l := strings.Index(a[i], "W")
		r := strings.LastIndex(a[i], "W")
		//fmt.Println(cur, l, r)
		if i&1 == 0 {
			if cur <= l {
				ans += r - cur
			} else {
				ans += cur - l + r - l
			}
			cur = r
		} else {
			if cur >= r {
				ans += cur - l
			} else {
				ans += r - cur + r - l
			}
			cur = l
		}
		//fmt.Println(ans)
	}

	fmt.Println(ans + last)
}
