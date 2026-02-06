package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var l, r int
	fmt.Fscan(in, &l, &r)
	fmt.Fprintln(out, (1<<bits.Len(uint(l^r)))-1)
}
