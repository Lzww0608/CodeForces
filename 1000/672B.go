package main

import "fmt"

func main() {
	var n int
	var s string
	fmt.Scan(&n)
	fmt.Scan(&s)
	if n > 26 {
		fmt.Println(-1)
		return
	}
	cnt := make([]int, 26)
	ans := 0
	for _, c := range s {
		x := int(c - 'a')
		if cnt[x] != 0 {
			ans++
		}
		cnt[x]++
	}
	fmt.Println(ans)
}
