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

	var t, n, k int
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &n, &k)
		if k > n || (k == 2 && n%2 == 1) || (n < k*2 && (n-k)%2 == 1) || (k%2 == 0 && n%k != 0 && n%k%2 == 1) {
			fmt.Fprintln(out, "NO")
			continue
		}
		fmt.Fprintln(out, "YES")
		m := n % k
		if m&1 == 0 {
			for i := 1; i < k; i++ {
				fmt.Fprint(out, n/k, " ")
			}
			fmt.Fprintln(out, m+n/k)
		} else {
			for i := 1; i < k; i++ {
				fmt.Fprint(out, n/k-1, " ")
			}
			fmt.Fprintln(out, n/k+m+k-1)
		}

	}

}
