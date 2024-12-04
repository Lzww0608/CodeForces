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

	var t int
	for fmt.Fscan(in, &t); t > 0; t-- {
		solve(in, out)
	}
}

func solve(in *bufio.Reader, out *bufio.Writer) {
	var n, m, k int
	fmt.Fscan(in, &n, &m, &k)
	a := make([]int, m)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}
	cnt := 0
	for i := range a {
		if a[i] > (n+k-1)/k {
			fmt.Fprintln(out, "NO")
			return
		} else if a[i] == (n+k-1)/k {
			cnt++
		}
	}
	if cnt <= (n-1)%k+1 {
		fmt.Fprintln(out, "YES")
	} else {
		fmt.Fprintln(out, "NO")
	}

}
