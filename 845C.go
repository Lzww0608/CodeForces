package main

import (
	"bufio"
	. "fmt"
	"os"
	"sort"
)

func main() {
	var n int
	in := bufio.NewReader(os.Stdin)
	Fscan(in, &n)

	type pair struct {
		x, y int
	}
	a := make([]pair, n)
	for i := range a {
		Fscan(in, &a[i].x, &a[i].y)
	}
	sort.Slice(a, func(i, j int) bool {
		if a[i].x == a[j].x {
			return a[i].y < a[j].y
		}
		return a[i].x < a[j].x
 	})

	end1, end2 := -1, -1
	for _, p := range a {
		if p.x <= end1 && p.x <= end2 {
			Println("NO")
			return
		} else if p.x > end1 {
			end1 = p.y
		} else {
			end2 = p.y
		}
	}
	Println("YES")
	return
}
