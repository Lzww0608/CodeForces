package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	ans := 0
	a := make([]int, m*2)
	for i := 0; i < n; i++ {
		for j := range a {
			fmt.Scan(&a[j])
			if j&1 == 1 && (a[j] == 1 || a[j-1] == 1) {
				ans++
			}

		}

	}
	fmt.Println(ans)
}
