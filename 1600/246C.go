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

	var n, k int
	fmt.Fscan(in, &n, &k)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	sort.Ints(a)
	i, j := 0, n
	for ; k > 0; k-- {
		fmt.Fprint(out, n-j+1, " ")
		for t := j; t < n; t++ {
			fmt.Fprint(out, a[t], " ")
		}
		fmt.Fprintln(out, a[i])
		if i++; i == j {
			j--
			i = 0
		}
	}
}
