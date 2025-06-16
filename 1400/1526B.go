package main

import (
	"bufio"
	"fmt"
	"os"
)

const N int = 11*111 - 11 - 111

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	for fmt.Fscan(in, &t); t > 0; t-- {
		solve(in, out)
	}
}

func solve(in *bufio.Reader, out *bufio.Writer) {
	var x int
	fmt.Fscan(in, &x)
	if x > N {
		fmt.Fprintln(out, "YES")
		return
	} else if x == N {
		fmt.Fprintln(out, "NO")
		return
	}

	for i := 0; i*111 <= x; i++ {
		for j := 0; i*111+j*11 <= x; j++ {
			y := i*111 + j*11
			if y > 0 && y%x == 0 {
				fmt.Fprintln(out, "YES")
				return
			}
		}
	}
	fmt.Fprintln(out, "NO")

}
