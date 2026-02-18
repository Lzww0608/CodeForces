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

	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	sort.Ints(a)

	ans := 0
	for i, j := 0, (n+1)/2; j < n; j++ {
		if a[j] >= a[i]*2 {
			ans++
			i++
		}
	}

	fmt.Fprintln(out, n-ans)
}
