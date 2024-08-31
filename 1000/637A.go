package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	cnt := map[int]int{}
	a := make([]int, n)
	mx := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
		cnt[a[i]]++
		mx = max(mx, cnt[a[i]])
	}

	clear(cnt)
	for _, x := range a {
		cnt[x]++
		if cnt[x] == mx {
			fmt.Println(x)
			break
		}
	}

	return
}
