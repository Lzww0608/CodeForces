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

	var n, k int
	fmt.Fscan(in, &n, &k)
	if k > (n-1)/2 {
		fmt.Fprintln(out, -1)
		return
	}
	vis := make([][]bool, n+1)
	for i := range vis {
		vis[i] = make([]bool, n+1)
	}
	fmt.Fprintln(out, k*n)
	for i := 1; i <= n; i++ {
		t := k
		for j := 1; j <= n && t > 0; j++ {
			if vis[i][j] || i == j {
				continue
			}
			vis[i][j] = true
			vis[j][i] = true
			fmt.Fprintln(out, i, j)
			t--
		}
	}

}
