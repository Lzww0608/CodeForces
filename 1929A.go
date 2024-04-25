package main

import (
	"bufio"
	. "fmt"
	"os"
	"sort"
)

func main() {

	in := bufio.NewReader(os.Stdin)
	//in := bufio.NewScanner(os.Stdin)

	var t, n int64
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &n)
		a := make([]int64, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		sort.Slice(a, func(i, j int) bool {
			return a[i] < a[j]
		})
		var ans int64
		for i := int64(1); i < n; i++ {
			ans += a[i] - a[i-1]
		}
		Println(ans)
	}

	return
}
