package main

import "fmt"

func main() {
	var t int
next:
	for fmt.Scan(&t); t > 0; t-- {
		var s string
		fmt.Scan(&s)
		a, b := "a"+s, s+"a"
		for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
			if a[i] != a[j] {
				fmt.Println("YES")
				fmt.Println(a)
				continue next
			}
			if b[i] != b[j] {
				fmt.Println("YES")
				fmt.Println(b)
				continue next
			}
		}
		fmt.Println("NO")
	}
}
