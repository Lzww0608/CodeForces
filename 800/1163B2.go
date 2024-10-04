package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n)
	cnt := map[int]int{}
	freq := map[int]int{}

	ans := 0
	for i := range a {
		fmt.Fscan(in, &a[i])
		x := a[i]
		cnt[x]++
		freq[cnt[x]] += cnt[x]
		if v := freq[cnt[x]]; v >= i && v < n {
			ans = v
		}
	}
	fmt.Fprintln(out, ans+1)
}
