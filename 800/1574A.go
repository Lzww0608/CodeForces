package main

import "fmt"

func main() {
	var t, n int
	for fmt.Scan(&t); t > 0; t-- {
		fmt.Scan(&n)
		ans := make([]string, 0, n)
		path := []byte{}

		var dfs func(int, int)
		dfs = func(a, b int) {
			if a == 0 && b == 0 {
				ans = append(ans, string(path))
				return
			}
			if len(ans) == n {
				return
			} else if a >= b {
				path = append(path, '(')
				dfs(a-1, b)
				path = path[:len(path)-1]
			} else {
				if a > 0 {
					path = append(path, '(')
					dfs(a-1, b)
					path = path[:len(path)-1]
				}
				path = append(path, ')')
				dfs(a, b-1)
				path = path[:len(path)-1]
			}
		}
		dfs(n, n)
		for i := range ans {
			fmt.Println(ans[i])
		}
	}
}
