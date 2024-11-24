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

	var n, k int
	fmt.Fscan(in, &n, &k)
	if k > (n/2)*n+(n&1)*((n+1)/2) {
		fmt.Fprintln(out, "NO")
		return
	}
	fmt.Fprintln(out, "YES")
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if k > 0 && (i+j)&1 == 0 {
				fmt.Fprint(out, "L")
				k--
			} else {
				fmt.Fprint(out, "S")
			}
		}
		fmt.Fprintln(out)
	}
}
