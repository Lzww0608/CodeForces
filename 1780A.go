package main

import "fmt"

func main() {
	var t, n, x int
	for fmt.Scan(&t); t > 0; t-- {
		fmt.Scan(&n)
		a := []int{}
		b := []int{}
		for i := 0; i < n; i++ {
			fmt.Scan(&x)
			if x&1 == 0 && len(a) < 2 {
				a = append(a, i+1)
			} else if x&1 == 1 && len(b) < 3 {
				b = append(b, i+1)
			}

		}
		if len(b) >= 3 {
			fmt.Println("YES")
			fmt.Println(b[0], b[1], b[2])
		} else if len(b) > 0 && len(a) > 1 {
			fmt.Println("YES")
			fmt.Println(b[0], a[0], a[1])
		} else {
			fmt.Println("NO")
		}
	}
}
