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
	fmt.Fprint(out, 2, " ")
	for i := 2; i <= n; i++ {
		fmt.Fprint(out, i*i*i+2*i*i+1, " ")
	}
}
