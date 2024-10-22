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
	s := [9]string{}
	t := make([][]byte, 9)
	for i := range s {
		fmt.Fscan(in, &s[i])
		t[i] = []byte(s[i])
	}
	for i := 0; i < 9; i++ {
		if t[i][i/3+i%3*3] != '1' {
			t[i][i/3+i%3*3] = '1'
		} else {
			t[i][i/3+i%3*3] = '2'
		}
	}
	for i := range t {
		fmt.Fprintln(out, string(t[i]))
	}

}
