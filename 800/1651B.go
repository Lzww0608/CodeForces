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

	var t, n int
	mx := 0
	for k := 1; k <= int(1e9); k *= 3 {
		mx++
	}

	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &n)
		if n > mx {
			fmt.Fprintln(out, "NO")
			continue
		}
		fmt.Fprintln(out, "YES")
		for i, k := 0, 1; i < n; i++ {
			fmt.Fprint(out, k, " ")
			k *= 3
		}
		fmt.Fprintln(out)
	}
	//fmt.Fprintln(out, mx)

}
