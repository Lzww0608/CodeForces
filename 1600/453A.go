package main

import "fmt"

func main() {
	var m float64
	var n int
	fmt.Scan(&m, &n)
	var ans, now, pre float64
	for i := float64(1); i <= m; i++ {
		now = quickPow(float64(i)/m, n)
		ans += (now - pre) * float64(i)
		pre = now
	}
	fmt.Printf("%.12f\n", ans)
}

func quickPow(x float64, n int) float64 {
	var res float64 = 1
	for n > 0 {
		if n&1 == 1 {
			res = res * x
		}
		x *= x
		n >>= 1
	}
	return res
}
