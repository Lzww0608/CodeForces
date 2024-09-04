package main

import "fmt"

func main() {
	var n, m int

	fmt.Scan(&n, &m)
	ans := 0
	for i := 0; i < 5; i++ {
		a := n / 5
		b := m / 5
		if i > 0 && (n%5) >= i {
			a++
		}
		if i > 0 && (m%5) >= 5-i {
			b++
		}
		ans += a * b
	}

	fmt.Println(ans)
}
