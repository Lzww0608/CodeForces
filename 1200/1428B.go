package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var dirs [][]int

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
	var s string
	fmt.Fscan(in, &s)

	if strings.Count(s, "<") == 0 || strings.Count(s, ">") == 0 {
		fmt.Fprintln(out, n)
		return
	}

	ans := 0
	for i := 0; i < n; i++ {
		if s[i] == '-' || s[(i-1+n)%n] == '-' {
			ans++
		}
	}
	fmt.Fprintln(out, ans)
}
