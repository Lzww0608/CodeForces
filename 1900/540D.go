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

	var R, S, P int
	fmt.Fscan(in, &R, &S, &P)
	f := make([][][]float64, R+1)
	for i := range f {
		f[i] = make([][]float64, S+1)
		for j := range f[i] {
			f[i][j] = make([]float64, P+1)
		}
	}
	f[R][S][P] = 1.0

	for sum := R + S + P; sum > 0; sum-- {
		for r := R; r >= 0; r-- {
			for s := S; s >= 0; s-- {
				p := sum - r - s
				if p < 0 || p > P {
					continue
				}
				all := float64(r*s + r*p + p*s)
				x := f[r][s][p]
				if r > 0 && p > 0 {
					f[r-1][s][p] += x * float64(r*p) / all
				}
				if s > 0 && r > 0 {
					f[r][s-1][p] += x * float64(s*r) / all
				}
				if p > 0 && s > 0 {
					f[r][s][p-1] += x * float64(s*p) / all
				}
			}
		}
	}

	a, b, c := 0.0, 0.0, 0.0
	for i := 1; i <= R; i++ {
		a += f[i][0][0]
	}
	for i := 1; i <= S; i++ {
		b += f[0][i][0]
	}
	for i := 1; i <= P; i++ {
		c += f[0][0][i]
	}

	fmt.Fprintln(out, a, b, c)
}
