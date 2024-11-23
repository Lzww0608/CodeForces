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

	var n, k int
	fmt.Fscan(in, &n, &k)
	for i := 1; i <= k; i++ {
		fmt.Fprint(out, i*2, i*2-1, " ")
	}
	for i := k + 1; i <= n; i++ {
		fmt.Fprint(out, i*2-1, i*2, " ")
	}
}
