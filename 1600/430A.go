package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const N int = 101

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)
	a := make([]int, n)
	id := make([]int, n)
	for i := range a {
		id[i] = i
		fmt.Fscan(in, &a[i])
	}
	b := make([][2]int, m)
	for i := range b {
		fmt.Fscan(in, &b[i][0], &b[i][1])
	}

	sort.Slice(id, func(i, j int) bool {
		return a[id[i]] < a[id[j]]
	})

	dif := make([]int, N)
	for j, i := range id {
		x := a[i]
		if j&1 == 0 {
			dif[x]++
		} else {
			dif[x]--
		}
	}
	for _, v := range b {
		l, r := v[0], v[1]
		sum := 0
		for _, x := range dif[l : r+1] {
			sum += x
		}
		if sum > 1 || sum < -1 {
			fmt.Fprintln(out, -1)
			return
		}
	}

	ans := make([]int, n)
	for j, i := range id {
		if j&1 == 0 {
			ans[i] = 1
		} else {
			ans[i] = 0
		}
	}
	for _, x := range ans {
		fmt.Fprint(out, x, " ")
	}

}
