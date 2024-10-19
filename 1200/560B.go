package main

import "fmt"

func main() {
	a := make([][2]int, 3)
	for i := range a {
		fmt.Scan(&a[i][0], &a[i][1])
	}
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			d1, h1 := a[1][i]+a[2][j], max(a[1][1-i], a[2][1-j])
			if d1 <= a[0][0] && h1 <= a[0][1] || d1 <= a[0][1] && h1 <= a[0][0] {
				fmt.Println("YES")
				return
			}
		}
	}
	fmt.Println("NO")
}
