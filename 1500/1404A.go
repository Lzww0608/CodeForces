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
	var n, k int
	fmt.Fscan(in, &n, &k)
	var s []byte
	fmt.Fscan(in, &s)

	for i := 0; i < n; i++ {
		if s[i] == '?' || s[i] == s[i%k] {
			continue
		} else if s[i%k] == '?' {
			s[i%k] = s[i]
		} else {
			fmt.Fprintln(out, "NO")
			return
		}
	}

	a, b := 0, 0
	for i := 0; i < k; i++ {
		if s[i] == '0' {
			a++
		} else if s[i] == '1' {
			b++
		}
	}
	if a > k/2 || b > k/2 {
		fmt.Fprintln(out, "NO")
		return
	}

	fmt.Fprintln(out, "YES")
}
