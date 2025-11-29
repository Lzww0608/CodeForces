package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	var s string
	fmt.Fscan(in, &s)

	blue := []int{}
	red := []int{}
	for i := range n {
		if s[i] == 'B' {
			blue = append(blue, a[i])
		} else {
			red = append(red, a[i])
		}
	}
	sort.Ints(blue)
	sort.Ints(red)

	for i, x := range blue {
		if x < i+1 {
			fmt.Fprintln(out, "NO")
			return
		}
	}

	t := len(blue) + 1
	for i, x := range red {
		if x > t+i {
			fmt.Fprintln(out, "NO")
			return
		}
	}

	fmt.Fprintln(out, "YES")
}
