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
	cnt := 1
	for cnt < n && n%cnt == 0 {
		cnt++
	}
	for i := 0; i < n; i++ {
		fmt.Fprint(out, string('a'+i%cnt))
	}
	fmt.Fprintln(out)
}
