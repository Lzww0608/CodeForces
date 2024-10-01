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

	var t int
	for fmt.Fscan(in, &t); t > 0; t-- {
		solve(in, out)
	}

}

func solve(in *bufio.Reader, out *bufio.Writer) {
	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n)
	f := false
	for i := range a {
		fmt.Fscan(in, &a[i])
		if a[i] == 1 {
			f = true
		}
	}
	sort.Ints(a)
	if a[0] == 0 && f {
		fmt.Fprintln(out, "NO")
	} else if f {
		for i := 0; i < n-1; i++ {
			if a[i]+1 == a[i+1] {
				fmt.Fprintln(out, "NO")
				return
			}
		}
		fmt.Fprintln(out, "YES")
	} else {
		fmt.Fprintln(out, "YES")
	}
}
