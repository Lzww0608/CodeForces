package main

import (
	"bufio"
	"fmt"
	"os"
)

const N int = 2e9

var square []int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, d int
	fmt.Fscan(in, &n, &d)
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}

	ans := 0
	for i, j := 0, 0; i < n; i++ {
		for j < n && a[j]-a[i] <= d {
			j++
		}
		ans += (j - i - 1) * (j - i - 2) / 2
	}

	fmt.Fprintln(out, ans)
}
