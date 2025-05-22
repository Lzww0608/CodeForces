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
	a := make([]int, 3)
	b := make([]int, 3)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}
	for i := range b {
		fmt.Fscan(in, &b[i])
	}

	mx := min(a[0], b[1]) + min(a[1], b[2]) + min(a[2], b[0])
	mn := max(0, a[0]-b[0]-b[2], a[1]-b[0]-b[1], a[2]-b[1]-b[2])
	fmt.Fprintln(out, mn, mx)
}
