package main

import "fmt"

func main() {
	var n, p int
	var s string
	fmt.Scan(&n, &p, &s)

	for i := 0; i < p; i++ {
		pre := s[i]
		for j := i + p; j < n; j += p {
			if s[j] == '.' || s[j] != pre {
				ans := []byte(s)
				if s[j] == '.' {
					if s[i] == '.' {
						ans[i], ans[j] = '0', '1'
					} else if s[i] == '0' {
						ans[j] = '1'
					} else {
						ans[j] = '0'
					}
				} else if s[i] == '.' {
					if s[j] == '0' {
						ans[i] = '1'
					} else {
						ans[i] = '0'
					}
				}
				for i := range ans {
					if ans[i] == '.' {
						ans[i] = '0'
					}
				}
				fmt.Println(string(ans))
				return
			}
		}
	}

	fmt.Println("No")
}
