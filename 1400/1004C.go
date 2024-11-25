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
	exist := make(map[int]bool)
	cnt := make(map[int]int)
	for i := range a {
		fmt.Fscan(in, &a[i])
		exist[a[i]] = true
		cnt[a[i]]++
	}

	ans := 0
	for _, x := range a[:n-1] {
		if cnt[x]--; cnt[x] == 0 {
			delete(cnt, x)
		}
		if _, ok := exist[x]; ok {
			delete(exist, x)
			ans += len(cnt)
			//fmt.Fprintln(out, cnt, ans)
		}
	}

	fmt.Fprintln(out, ans)
}
