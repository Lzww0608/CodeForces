package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)
	if n&1 == 0 {
		fmt.Println("NO")
		return
	}
	fmt.Println("YES")
	a := make([]int, n*2)
	for i := 0; i < n; i++ {
		if i&1 == 0 {
			a[i], a[i+n] = i*2+1, i*2+2
		} else {
			a[i], a[i+n] = i*2+2, i*2+1
		}

	}
	for _, x := range a {
		fmt.Print(x, " ")
	}
	fmt.Println()
}
