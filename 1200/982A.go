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

	var n int
	fmt.Fscan(in, &n)
	var s string
	fmt.Fscan(in, &s)
	s = "0" + s + "0"
	if strings.Contains(s, "000") || strings.Contains(s, "11") {
		fmt.Fprintln(out, "No")
	} else {
		fmt.Fprintln(out, "Yes")
	}
}
