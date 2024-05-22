package main

import "fmt"

func main() {
	var t, n int
	for fmt.Scan(&t); t > 0; t-- {
		fmt.Scan(&n)
		a := make([]int, n)
		for i := range a {
			fmt.Scan(&a[i])

		}
		i := 0
		for i < n {
			if a[i] != i+1 {
				break
			}
			i++
		}
		if i < n {
			j := i + 1
			for j < n && a[j] != i+1 {
				j++
			}
			for i < j {
				a[i], a[j] = a[j], a[i]
				i++
				j--
			}
		}
		for _, x := range a {
			fmt.Print(x, " ")
		}
		fmt.Println()
	}
}
