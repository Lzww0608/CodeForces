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

	var n, m int
	fmt.Fscan(in, &n, &m)
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}

	sort.Ints(a)
	ans := 0
	s := make([]int, m)
	for i, x := range a {
		s[i%m] += x
		ans += s[i%m]
		fmt.Fprintf(out, "%d ", ans)
	}

}
