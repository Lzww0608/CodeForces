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

	var h int
	fmt.Fscan(in, &h)
	a := make([]int, h+1)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}
	
	f := true
	for i := 1; i <= h; i++ {
		if a[i] > 1 && a[i-1] > 1 {
			f = false
			break
		}
	}
	if f {
		fmt.Fprintln(out, "perfect")
		return
	}

	fmt.Fprintln(out, "ambiguous")
	cur, pre := 1, 0
	for _, x := range a {
		next := cur
		for j := 0; j < x; j++ {
			fmt.Fprint(out, pre, " ")
		}
		pre = next
		cur += x
	}
	fmt.Fprintln(out)

	cur = 1
	last := []int{0}
	for _, x := range a {
		tmp := []int{}
		for j := 0; j < x; j++ {
			fmt.Fprint(out, last[j%len(last)], " ")
			tmp = append(tmp, cur+j)
		}
		cur += x
		last = tmp
	}
	fmt.Fprintln(out)
}
