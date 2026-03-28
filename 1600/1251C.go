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
	var a string
	fmt.Fscan(in, &a)
	n := len(a)
	odd := make([]byte, 0, n)
	even := make([]byte, 0, n)
	for i := range a {
		x := int(a[i] - '0')
		if x&1 == 1 {
			odd = append(odd, a[i])
		} else {
			even = append(even, a[i])
		}
	}

	ans := make([]byte, 0, n)
	for len(odd) > 0 || len(even) > 0 {
		if len(odd) == 0 || (len(even) > 0 && even[0] < odd[0]) {
			ans = append(ans, even[0])
			even = even[1:]
		} else {
			ans = append(ans, odd[0])
			odd = odd[1:]
		}
	}

	fmt.Fprintln(out, string(ans))
}
