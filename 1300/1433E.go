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
	F := [21]int{0, 0, 1, 3, 3, 15, 40, 280, 1260, 11340, 72576, 798336,
		6652800, 86486400, 889574400, 13343616000, 163459296000, 2778808032000, 39520825344000,
		750895681536000, 12164510040883200}
	fmt.Fprintln(out, F[n])
}
