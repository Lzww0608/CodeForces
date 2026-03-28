package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	var l, r int
	fmt.Fscan(in, &l, &r)
	a, b := strconv.Itoa(l), strconv.Itoa(r)
	n := len(a)
	if l == r {
		fmt.Fprintln(out, n*2)
		return
	}

	p := 0
	for p < n && a[p] == b[p] {
		p++
	}

	ans := p * 2
	if a[p]+1 < b[p] {
		fmt.Fprintln(out, ans)
		return
	}

	ans++
	p++
	for p < n && a[p] == '9' && b[p] == '0' {
		p++
		ans++
	}
	fmt.Fprintln(out, ans)

}
