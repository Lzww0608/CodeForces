package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var k, n int
	fmt.Fscan(in, &k, &n)
	a := make([]int, k)
	vis := make([]bool, k+1)
	pre := make([]int, k+1)
	for i := 0; i < k; i++ {
		fmt.Fscan(in, &a[i])
		pre[i+1] = pre[i] + a[i]
	}
	sort.Ints(pre[1:])

	for i := 1; i <= k; i++ {
		if pre[i] == pre[i-1] {
			vis[i] = true
		}
	}

	var x int
	cnt := make(map[int]int)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &x)
		for j := 1; j <= k; j++ {
			if vis[j] {
				continue
			}
			cnt[x-pre[j]]++
		}
	}

	ans := 0
	for _, y := range cnt {
		if y == n {
			ans++
		}
	}
	fmt.Fprintln(out, ans)
}
