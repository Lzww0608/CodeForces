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

		var t int
		for fmt.Fscan(in, &t); t > 0; t-- {
			solve(in, out)
		}
	}

	func solve(in *bufio.Reader, out *bufio.Writer) {
		var n int
		fmt.Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			fmt.Fscan(in, &a[i])
		}

		vis := make([]bool, n)
		for i, x := range a {
			y := ((i+x)%n + n) % n
			if vis[y] {
				fmt.Fprintln(out, "NO")
				return
			}
			vis[y] = true
		}

		fmt.Fprintln(out, "YES")
	}