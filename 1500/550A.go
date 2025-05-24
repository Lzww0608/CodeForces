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

	var s string
	fmt.Fscan(in, &s)

	pos := strings.Index(s, "AB")
	if pos == -1 {
		fmt.Fprintln(out, "NO")
		return
	}

	if strings.Contains(s[pos+2:], "BA") || strings.Contains(s[:pos], "BA") {
		fmt.Fprintln(out, "YES")
		return
	}

	pos = strings.Index(s, "BA")
	if pos == -1 {
		fmt.Fprintln(out, "NO")
		return
	}

	if strings.Contains(s[pos+2:], "AB") || strings.Contains(s[:pos], "AB") {
		fmt.Fprintln(out, "YES")
		return
	}

	fmt.Fprintln(out, "NO")
	return
}
