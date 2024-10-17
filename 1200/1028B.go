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

	var n, m int
	fmt.Fscan(in, &n, &m)
	x := 5
	cnt := 0
	for x < n {
		x += 9
		cnt++
	}
	fmt.Fprint(out, strings.Repeat("45", cnt))
	fmt.Fprintln(out, "5")
	fmt.Fprint(out, strings.Repeat("54", cnt))
	fmt.Fprintln(out, "5")
}
