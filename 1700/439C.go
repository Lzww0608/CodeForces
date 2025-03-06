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

	var n, k, p, x int
	fmt.Fscan(in, &n, &k, &p)
	odd, even := make([]int, 0, n), make([]int, 0, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &x)
		if x&1 == 0 {
			even = append(even, x)
		} else {
			odd = append(odd, x)
		}
	}

	if len(odd) < k-p || (len(odd)-(k-p))&1 == 1 || len(even)+(len(odd)-(k-p))/2 < p {
		fmt.Fprintln(out, "NO")
		return
	}

	fmt.Fprintln(out, "YES")
	for i := 0; i < p; i++ {
		if i == k-1 {
			fmt.Fprint(out, len(even)+len(odd), " ")
			for _, x := range even {
				fmt.Fprint(out, x, " ")
			}
			for _, x := range odd {
				fmt.Fprint(out, x, " ")
			}
			return
		}

		if len(even) > 0 {
			fmt.Fprintln(out, 1, even[0])
			even = even[1:]
		} else {
			fmt.Fprintln(out, 2, odd[0], odd[1])
			odd = odd[2:]
		}
	}

	for i := 0; i < k-p; i++ {
		if i == k-p-1 {
			fmt.Fprint(out, len(even)+len(odd), " ")
			for _, x := range even {
				fmt.Fprint(out, x, " ")
			}
			for _, x := range odd {
				fmt.Fprint(out, x, " ")
			}
		} else {
			fmt.Fprintln(out, 1, odd[0])
			odd = odd[1:]
		}
	}
}
