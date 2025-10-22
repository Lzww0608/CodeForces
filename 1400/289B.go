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

	var n, m, d int
	fmt.Fscan(in, &n, &m, &d)
	a := make([]int, n*m)
	for i := range a {
		fmt.Fscan(in, &a[i])
		if (a[i]-a[0])%d != 0 {
			fmt.Fprintln(out, -1)
			return
		}
	}
	sort.Ints(a)
	ans := 0
	for _, x := range a {
		ans += abs(x-a[(n*m)/2]) / d
	}

	fmt.Fprintln(out, ans)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
