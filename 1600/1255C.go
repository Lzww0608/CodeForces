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

	var a, b, c int
	cnt := make(map[int]int)
	rev := make(map[[2]int]int)
	two := make(map[[2]int][]int)
	for i := 0; i < n-2; i++ {
		fmt.Fscan(in, &a, &b, &c)
		cnt[a]++
		cnt[b]++
		cnt[c]++
		rev[[2]int{c, b}] = c
		two[[2]int{a, b}] = append(two[[2]int{a, b}], c)
		two[[2]int{b, a}] = append(two[[2]int{b, a}], c)
		two[[2]int{a, c}] = append(two[[2]int{a, c}], b)
		two[[2]int{c, a}] = append(two[[2]int{c, a}], b)
		two[[2]int{b, c}] = append(two[[2]int{b, c}], a)
		two[[2]int{c, b}] = append(two[[2]int{c, b}], a)
	}

	vis := make(map[int]bool)
	ans := make([]int, n)
	for k, v := range cnt {
		if v == 1 {
			vis[k] = true
			if ans[0] == 0 {
				ans[0] = k
			} else {
				ans[n-1] = k
			}
		}
	}
	//fmt.Fprintln(out, ans)
	for k, v := range cnt {
		if v == 2 {
			vis[k] = true
			if _, ok := two[[2]int{ans[0], k}]; ans[n-2] != 0 || ok && ans[1] == 0 {
				ans[1] = k
			} else {
				ans[n-2] = k
			}
		}
	}

	for i := 2; i < n-2; i++ {
		a, b := ans[i-2], ans[i-1]
		for _, v := range two[[2]int{a, b}] {
			if !vis[v] {
				ans[i] = v
				break
			}
		}
		vis[ans[i]] = true
	}

	for _, x := range ans {
		fmt.Fprint(out, x, " ")
	}
}
