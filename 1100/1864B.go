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

	var t, n, k int

	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &n, &k)
		var s string
		fmt.Fscan(in, &s)
		t := []byte(s)
		if k&1 == 0 {
			sort.Slice(t, func(i, j int) bool { return t[i] < t[j] })
		} else {
			for i := 0; i < 2; i++ {
				tmp := []byte{}
				for j := i; j < n; j += 2 {
					tmp = append(tmp, t[j])
				}
				sort.Slice(tmp, func(i, j int) bool { return tmp[i] < tmp[j] })
				for j := i; j < n; j += 2 {
					t[j] = tmp[j/2]
				}
			}
		}

		fmt.Fprintln(out, string(t))
	}
}
