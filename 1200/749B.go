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

	a := make([][2]int, 3)
	for i := 0; i < 3; i++ {
		fmt.Fscan(in, &a[i][0], &a[i][1])
	}
	fmt.Fprintln(out, 3)
	x, y := a[2][0]-a[1][0], a[2][1]-a[1][1]
	fmt.Fprintln(out, a[0][0]+x, a[0][1]+y)
	fmt.Fprintln(out, a[0][0]-x, a[0][1]-y)
	x, y = a[1][0]-a[0][0], a[1][1]-a[0][1]
	fmt.Fprintln(out, a[2][0]+x, a[2][1]+y)
}
