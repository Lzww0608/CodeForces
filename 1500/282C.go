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

	var s, t string
	fmt.Fscan(in, &s, &t)
	if len(s) != len(t) || !strings.Contains(s, "1") && strings.Contains(t, "1") {
		fmt.Fprintln(out, "NO")
	} else if strings.Contains(s, "1") && !strings.Contains(t, "1") {
		fmt.Fprintln(out, "NO")
	} else {
		fmt.Fprintln(out, "YES")
	}

}
