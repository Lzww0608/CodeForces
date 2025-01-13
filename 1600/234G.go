package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fin, _ := os.Open("input.txt")
	fout, _ := os.Create("output.txt")
	in := bufio.NewReader(fin)
	out := bufio.NewWriter(fout)
	defer fin.Close()
	defer fout.Close()
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)
	a, b := 1, 0
	for a < n {
		a <<= 1
		b++
	}

	fmt.Fprintln(out, b)
	for i := 0; i < b; i++ {
		c := []int{}
		for j := 0; j < n; j++ {
			if j&(1<<i) != 0 {
				c = append(c, j+1)
			}
		}
		fmt.Fprint(out, len(c), " ")
		for _, v := range c {
			fmt.Fprint(out, v, " ")
		}
		fmt.Fprintln(out)
	}
}
