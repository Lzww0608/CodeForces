package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	var n, m int
	fmt.Fscan(in, &n, &m)
	a := make([]int, 0, n)
	cnt := 0
	for i := 0; i < n; i++ {
		var x int
		fmt.Fscan(in, &x)
		cnt += len(strconv.Itoa(x))
		t := 0
		for x > 0 && x%10 == 0 {
			t++
			x /= 10
		}
		if t > 0 {
			a = append(a, t)
		}
	}

	sort.Slice(a, func(i, j int) bool { return a[i] > a[j] })
	for i := 0; i < len(a); i += 2 {
		cnt -= a[i]
	}

	if cnt-1 >= m {
		fmt.Fprintln(out, "Sasha")
	} else {
		fmt.Fprintln(out, "Anna")
	}
}
