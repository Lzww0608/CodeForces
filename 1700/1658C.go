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
		sovle(in, out)
	}
}

func sovle(in *bufio.Reader, out *bufio.Writer) {
	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n)
	cnt, pos := 0, 0
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
		if a[i] == 1 {
			cnt++
			pos = i
		}
	}
	if cnt != 1 {
		fmt.Fprintln(out, "NO")
		return
	}
	for i := 1; i < n; i++ {
		if a[(i+pos)%n]-a[(i+pos-1+n)%n] > 1 {
			fmt.Fprintln(out, "NO")
			return
		}
	}
	fmt.Fprintln(out, "YES")
	return
}
