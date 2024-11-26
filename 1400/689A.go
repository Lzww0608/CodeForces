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

	var n int
	fmt.Fscan(in, &n)
	var s string
	fmt.Fscan(in, &s)
	m := map[int]bool{}
	for i := 0; i < n; i++ {
		m[int(s[i]-'0')] = true
	}
	if (m[1] || m[2] || m[3]) && m[0] {
		fmt.Fprintln(out, "YES")
		return
	}

	if (m[1] || m[2] || m[3]) && (m[7] || m[9]) && (m[1] || m[4] || m[7]) && (m[3] || m[6] || m[9]) {
		fmt.Fprintln(out, "YES")
		return
	}

	fmt.Fprintln(out, "NO")
}
