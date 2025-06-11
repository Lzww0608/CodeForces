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
	var n int
	fmt.Fscan(in, &n)
	m := make(map[int]bool)
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			m[i] = true
			n /= i
			break
		}
	}

	for i := 3; i*i <= n; i++ {
		if n%i == 0 && !m[i] {
			m[i], m[n/i] = true, true
			break
		}
	}

	if len(m) != 3 || m[1] {
		fmt.Fprintln(out, "NO")
	} else {
		fmt.Fprintln(out, "YES")
		for k := range m {
			fmt.Fprintf(out, "%d ", k)
		}
		fmt.Fprintln(out)
	}
}
