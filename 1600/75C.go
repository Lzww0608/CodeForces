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

	var a, b int
	fmt.Fscan(in, &a, &b)
	x := gcd(a, b)
	c := []int{}
	for i := 1; i*i <= x; i++ {
		if x%i == 0 {
			c = append(c, i)
			if i*i != x {
				c = append(c, x/i)
			}
		}
	}
	sort.Ints(c)

	var n int
	fmt.Fscan(in, &n)

	for range n {
		var l, r int
		fmt.Fscan(in, &l, &r)
		if c[len(c)-1] < l {
			fmt.Fprintln(out, -1)
			continue
		}
		p := sort.SearchInts(c, r+1)
		if p-1 >= 0 && c[p-1] >= l {
			fmt.Fprintln(out, c[p-1])
		} else {
			fmt.Fprintln(out, -1)
		}
	}
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}
