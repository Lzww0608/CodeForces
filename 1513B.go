package main

import (
	"bufio"
	. "fmt"
	"os"
)

var mod int = 1e9 + 7

func main() {
	var t int
	in := bufio.NewReader(os.Stdin)
	for Fscan(in, &t); t > 0; t-- {
		var n int
		Fscan(in, &n)
		a := make([]int, n)
		x := -1
		for i := range a {
			Fscan(in, &a[i])
			x &= a[i]
		}
		cnt := 0
		for _, num := range a {
			if num == x {
				cnt++
			}
		}
		ans := cnt * (cnt - 1) % mod
		for i := 2; i <= n-2; i++ {
			ans = ans * i % mod
		}
		Println(ans)
	}
	return
}
