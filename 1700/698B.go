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
	a := make([]int, n+1)
	vis := make([]int, n+1)
	root := 0
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
		if i == a[i] {
			root = i
		}
	}
	cnt := 0
	ans := 0
	for i := 1; i <= n; i++ {
		if vis[i] != 0 {
			continue
		}
		cnt++
		j := i
		vis[j] = cnt
		for vis[a[j]] == 0 {
			j = a[j]
			vis[j] = cnt
		}

		if vis[a[j]] == cnt {
			if root == 0 {
				root = j
			}

			if a[j] != root {
				ans++
				a[j] = root
			}
		}
	}

	fmt.Fprintln(out, ans)
	for i := 1; i <= n; i++ {
		fmt.Fprint(out, a[i], " ")
	}
}
