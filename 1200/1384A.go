package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
	for i := range a {
		fmt.Fscan(in, &a[i])
	}
	s := []byte(strings.Repeat("a", 51))
	fmt.Fprintln(out, string(s))
	for _, x := range a {
		s[x] = (s[x]-'a'+1)%26 + 'a'
		fmt.Fprintln(out, string(s))
	}
}
