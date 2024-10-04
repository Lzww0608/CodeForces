package main

import "fmt"

func main() {
	var t, n, k1, k2, w, b int
	for fmt.Scan(&t); t > 0; t-- {
		fmt.Scan(&n, &k1, &k2)
		fmt.Scan(&w, &b)
		if (k1+k2)/2 < w || (n*2-k1-k2)/2 < b {
			fmt.Println("NO")
		} else {
			fmt.Println("YES")
		}

	}
}
