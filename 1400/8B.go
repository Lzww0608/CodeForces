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

	var s string
	fmt.Fscan(in, &s)
	x, y := 0, 0
	type pair struct {
		x, y int
	}
	vis := make(map[pair]bool)
	vis[pair{0, 0}] = true
	for i := range s {
		if s[i] == 'R' {
			x++
		} else if s[i] == 'L' {
			x--
		} else if s[i] == 'U' {
			y++
		} else if s[i] == 'D' {
			y--
		}
		cnt := 0
		if vis[pair{x + 1, y}] {
			cnt++
		}
		if vis[pair{x - 1, y}] {
			cnt++
		}
		if vis[pair{x, y + 1}] {
			cnt++
		}
		if vis[pair{x, y - 1}] {
			cnt++
		}
		if cnt > 1 || vis[pair{x, y}] {
			fmt.Fprintln(out, "BUG")
			return
		}
		vis[pair{x, y}] = true
	}

	fmt.Fprintln(out, "OK")
}
