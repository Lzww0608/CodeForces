package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
)

const N int = 2e9

var square []int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	for i := 3; i*i < N; i += 2 {
		square = append(square, i*i)
	}

	var t int
	for fmt.Fscan(in, &t); t > 0; t-- {
		solve(in, out)
	}
}

func solve(in io.Reader, out io.Writer) {
	var n int
	fmt.Fscan(in, &n)

	// a ^ 2 = b * 2 + 1
	// b <= n - 1, b * 2 + 1 <= 2 * n - 1
	i := sort.SearchInts(square, 2*n)
	fmt.Fprintln(out, i)
}
