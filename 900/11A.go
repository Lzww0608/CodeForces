package main

import "fmt"

func main() {
	var n, d int

	fmt.Scan(&n, &d)
	a := make([]int, n)
	mx := 0
	for i := range a {
		fmt.Scan(&a[i])
		a[i] -= i
		mx = max(mx, a[i])
	}
	pre, ans := 0, 0
	for _, x := range a {
		if x < pre {
			t := (pre - x + d - 1) / d
			ans += t
			pre = x + d*t
		} else {
			pre = x
		}
	}

	fmt.Println(ans)
}
