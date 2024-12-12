package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, k int
	fmt.Fscan(in, &n, &k)
	var s string
	a := make([]bool, n-k+1)
	for i := range a {
		fmt.Fscan(in, &s)
		if s == "NO" {
			a[i] = false
		} else {
			a[i] = true
		}
	}

	name := make([]string, 50)
	for i := 0; i < 5; i++ {
		for j := 0; j < 10; j++ {
			c := string(byte('A' + j))
			tmp := c + strings.Repeat(string(byte('a'+j)), i)
			name[i*10+j] = tmp
		}
	}

	ans := make([]string, n)
	for i := 0; i < n; i++ {
		if i >= k-1 && !a[i-k+1] {
			ans[i] = ans[i-k+1]
		} else {
			ans[i] = name[i]
		}
	}
	for _, v := range ans {
		fmt.Fprint(out, v, " ")
	}
}
