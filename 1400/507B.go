package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var r, x1, y1, x2, y2 int
	fmt.Fscan(in, &r, &x1, &y1, &x2, &y2)
	d := math.Sqrt(float64((x1-x2)*(x1-x2)) + float64((y1-y2)*(y1-y2)))
	ans := math.Ceil(d / (2 * float64(r)))
	fmt.Fprintln(out, ans)
}
