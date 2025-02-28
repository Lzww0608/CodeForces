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
	if n > k*(k-1) {
		fmt.Fprintln(out, "NO")
		return
	}
	fmt.Fprintln(out, "YES")
	cnt := 0
	for i := 0; i < k && cnt < n; i++ {
		for j := 1; j <= k && cnt < n; j++ {
			fmt.Fprintln(out, j, (i+j)%k+1)
			cnt++
		}
	}
}
