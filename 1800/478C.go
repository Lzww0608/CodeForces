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

	var r, g, b int
	fmt.Fscan(in, &r, &g, &b)
	mx := max(r, g, b)

	if mx >= (r+g+b-mx)*2 {
		fmt.Fprintln(out, r+g+b-mx)
	} else {
		fmt.Fprintln(out, (r+g+b)/3)
	}
}
