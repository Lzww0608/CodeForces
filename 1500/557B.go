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

	var n, w int
	fmt.Fscan(in, &n, &w)
	a := make([]int, n*2)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}
	sort.Ints(a)
	ans := min(float64(w)/float64(n*3), float64(a[0]), float64(a[n])/2) * 3 * float64(n)
	fmt.Fprintln(out, ans)
}
