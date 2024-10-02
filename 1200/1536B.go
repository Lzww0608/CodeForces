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

	var t int
	for fmt.Fscan(in, &t); t > 0; t-- {
		solve(in, out)
	}
}

func solve(in *bufio.Reader, out *bufio.Writer) {
	var n int
	var s string
	fmt.Fscan(in, &n)
	fmt.Fscan(in, &s)
	cnt := 3*n - 1
	t := []string{""}
	for j := 0; j < 3 && cnt > 0; j++ {
		sz := len(t)
		for i := 0; i < sz && cnt > 0; i++ {
			for k := 0; k < 26 && cnt > 0; k++ {
				t = append(t, t[i]+string('a'+k))
				cnt--
			}
		}
	}

	for i := range t {
		if !strings.Contains(s, t[i]) {
			fmt.Fprintln(out, t[i])
			return
		}
	}

}
