package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t, n int
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			fmt.Fscan(in, &a[i])
		}

		sort.Ints(a)
		if n == 2 {
			fmt.Fprintln(out, a[0], a[1])
			continue
		}
		idx := -1
		pos, mn := 0, math.MaxInt32
		for i := 0; i < n-1; i++ {
			if a[i] == a[i+1] {
				idx = i
				break
			}
			if a[i+1]-a[i] < mn {
				mn = a[i+1] - a[i]
				pos = i
			}
		}
		if idx == -1 {
			idx = pos
		}

		for i := idx + 1; i < n+idx+1; i++ {
			fmt.Fprint(out, a[i%n], " ")
		}
		fmt.Fprintln(out)
	}

}
