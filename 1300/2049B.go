package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const N int = 10_000_001

var isPrime [N]bool

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	for i := 2; i < N; i++ {
		if isPrime[i] {
			continue
		}

		for j := i * i; j < N; j += i {
			isPrime[j] = true
		}
	}

	var t int
	for fmt.Fscan(in, &t); t > 0; t-- {
		solve(in, out)
	}
}

func solve(in *bufio.Reader, out *bufio.Writer) {
	var n int
	fmt.Fscan(in, &n)
	var s string
	fmt.Fscan(in, &s)

	if !strings.Contains(s, "s") || !strings.Contains(s, "p") {
		fmt.Fprintln(out, "YES")
		return
	}

	a := strings.LastIndex(s, "s")
	b := strings.Index(s, "p")
	if a > b {
		fmt.Fprintln(out, "NO")
		return
	}

	if b < n - 1 && b > 0 && strings.Contains(s[1:b], "s") {
		fmt.Fprintln(out, "NO")
	} else {
		fmt.Fprintln(out, "YES")
	}
}
