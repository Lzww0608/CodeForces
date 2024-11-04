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
	var s string
	fmt.Fscan(in, &s)
	m := map[byte]int{}
	for i := range s {
		m[s[i]]++
	}

	if len(m) == 3 {
		fmt.Fprintln(out, "BGR")
	} else if len(m) == 1 {
		for k := range m {
			fmt.Fprintln(out, string(k))
		}
	} else {
		if m['R'] > 0 && m['G'] > 0 {
			if m['R'] == 1 {
				if m['G'] == 1 {
					fmt.Fprintln(out, "B")
				} else {
					fmt.Fprintln(out, "BR")
				}
			} else {
				if m['G'] == 1 {
					fmt.Fprintln(out, "BG")
				} else {
					fmt.Fprintln(out, "BGR")
				}
			}
			return
		}

		if m['R'] > 0 && m['B'] > 0 {
			if m['R'] == 1 {
				if m['B'] == 1 {
					fmt.Fprintln(out, "G")
				} else {
					fmt.Fprintln(out, "GR")
				}
			} else {
				if m['B'] == 1 {
					fmt.Fprintln(out, "BG")
				} else {
					fmt.Fprintln(out, "BGR")
				}
			}
			return
		}

		if m['B'] > 0 && m['G'] > 0 {
			if m['B'] == 1 {
				if m['G'] == 1 {
					fmt.Fprintln(out, "R")
				} else {
					fmt.Fprintln(out, "BR")
				}
			} else {
				if m['G'] == 1 {
					fmt.Fprintln(out, "GR")
				} else {
					fmt.Fprintln(out, "BGR")
				}
			}
			return
		}
	}
}
