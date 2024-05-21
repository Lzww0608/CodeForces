package main

import "fmt"

func main() {
	var t int
	for fmt.Scan(&t); t > 0; t-- {
		var s string
		fmt.Scan(&s)
		cnt := 0
		for i := range s {
			if s[i] == 'N' {
				cnt++
				if cnt > 1 {
					break
				}
			}
		}
		if cnt == 1 {
			fmt.Println("NO")
		} else {
			fmt.Println("YES")
		}

	}

}
