package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
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
	var n, k int
	fmt.Fscan(in, &n, &k)
	var str string
	fmt.Fscan(in, &str)

	s := []byte(str)
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	if k == n {
		fmt.Fprintln(out, string(s[n-1]))
		return
	}
	cnt := 0
	for i := 0; i < n; i++ {
		if s[i] != s[0] {
			break
		} else {
			cnt++
		}
	}

	if k > cnt {
		fmt.Fprintln(out, string(s[k-1]))
		return
	}
	f := true
	for i := k; i < n; i++ {
		if s[i] != s[k] {
			f = false
			break
		}
	}

	if f {
		fmt.Fprint(out, string(s[0])+strings.Repeat(string(s[k]), (n-k)/k))
		if (n-k)%k != 0 {
			fmt.Fprintln(out, string(s[k]))
		} else {
			fmt.Fprintln(out)
		}
	} else {
		fmt.Fprintln(out, string(s[0])+string(s[k:]))
	}

}
