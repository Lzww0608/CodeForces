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

	var n int
	fmt.Fscan(in, &n)
	s := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &s[i])
	}

	check1 := func() bool {
		for i := 0; i < n; i++ {
			f := true
			for j := 0; j < n; j++ {
				if s[i][j] != 'E' {
					f = false
					break
				}
			}
			if f {
				return false
			}
		}
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if s[i][j] != 'E' {
					fmt.Fprintln(out, i+1, j+1)
					break
				}
			}
		}
		return true
	}

	check2 := func() bool {
		for j := 0; j < n; j++ {
			f := true
			for i := 0; i < n; i++ {
				if s[i][j] != 'E' {
					f = false
					break
				}
			}
			if f {
				return false
			}
		}
		for j := 0; j < n; j++ {
			for i := 0; i < n; i++ {
				if s[i][j] != 'E' {
					fmt.Fprintln(out, i+1, j+1)
					break
				}
			}
		}
		return true
	}

	if !check1() && !check2() {
		fmt.Fprintln(out, -1)
		return
	}

}
