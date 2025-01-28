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

	var k int
	fmt.Fscan(in, &k)
	t := 1
	for t <= 100_000 {
		t <<= 1
	}
	// t + k, t,     0
	// k,     t + k, k
	fmt.Fprintln(out, 2, 3)
	fmt.Fprintln(out, t+k, t, 0)
	fmt.Fprintln(out, k, t+k, k)
}
