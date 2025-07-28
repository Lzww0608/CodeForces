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
	a := []int{1}
	for i := 2; i*i <= n; i++ {
		for p, d := i, a; n%i == 0; n /= i {
			for _, d := range d {
				a = append(a, d*p)
			}
			p *= i
		}
	}
	if n > 1 {
		for _, d := range a {
			a = append(a, d*n)
		}
	}
	sort.Ints(a)
	if k > len(a) {
		fmt.Fprint(out, -1)
	} else {
		fmt.Fprint(out, a[k-1])
	}

}
