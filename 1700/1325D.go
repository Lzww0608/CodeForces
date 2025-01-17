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

	var u, v int
	fmt.Fscan(in, &u, &v)
	if v < u || u%2 != v%2 {
		fmt.Fprintln(out, -1)
	} else if u == 0 && v == 0 {
		fmt.Fprintln(out, 0)
	} else if u == v {
		fmt.Fprintln(out, 1)
		fmt.Fprintln(out, u)
	} else {
		w := (v - u) / 2
		if u&w == 0 {
			fmt.Fprintln(out, 2)
			fmt.Fprintln(out, u+w, w)
		} else {
			fmt.Fprintln(out, 3)
			fmt.Fprintln(out, w, w, u)
		}
	}
	return
}
