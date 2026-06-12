package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
)

type pair struct {
	x, y int
	c    byte
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)
	a := make([][]int, n)
	f := make([][][2]pair, n)
	ans := math.MaxInt
	pos := [2]int{-1, -1}
	for i := range a {
		a[i] = make([]int, n)
		f[i] = make([][2]pair, n)
		for j := range a[i] {
			fmt.Fscan(in, &a[i][j])
			if a[i][j] == 0 {
				ans = 1
				pos = [2]int{i, j}
			}
			x, y := 0, 0
			t := a[i][j]
			for t != 0 && t%2 == 0 {
				t /= 2
				x++
			}
			for t != 0 && t%5 == 0 {
				t /= 5
				y++
			}

			f[i][j] = [2]pair{{x, y, ' '}, {x, y, ' '}}
		}
	}

	for i := range n {
		for j := range n {
			if i == 0 {
				if j == 0 {
					continue
				} else {
					v := f[i][j-1]
					f[i][j][0].x += v[0].x
					f[i][j][0].y += v[0].y
					f[i][j][1].x += v[1].x
					f[i][j][1].y += v[1].y
					f[i][j][0].c = 'R'
					f[i][j][1].c = 'R'
				}
			} else if j == 0 {
				v := f[i-1][j]
				f[i][j][0].x += v[0].x
				f[i][j][0].y += v[0].y
				f[i][j][1].x += v[1].x
				f[i][j][1].y += v[1].y
				f[i][j][0].c = 'D'
				f[i][j][1].c = 'D'
			} else {
				v1 := f[i][j-1]
				v2 := f[i-1][j]
				if v1[0].x < v2[0].x || (v1[0].x == v2[0].x && v1[0].y < v2[0].y) {
					f[i][j][0].x += v1[0].x
					f[i][j][0].y += v1[0].y
					f[i][j][0].c = 'R'
				} else {
					f[i][j][0].x += v2[0].x
					f[i][j][0].y += v2[0].y
					f[i][j][0].c = 'D'
				}

				if v1[1].y < v2[1].y || (v1[1].y == v2[1].y && v1[1].x < v2[1].x) {
					f[i][j][1].x += v1[1].x
					f[i][j][1].y += v1[1].y
					f[i][j][1].c = 'R'
				} else {
					f[i][j][1].x += v2[1].x
					f[i][j][1].y += v2[1].y
					f[i][j][1].c = 'D'
				}
			}
		}
	}

	buildPath := func(n, k int) string {
		path := make([]byte, 0, 2*n-2)
		i, j := n-1, n-1

		for i > 0 || j > 0 {
			c := f[i][j][k].c
			path = append(path, c)
			if c == 'R' {
				j--
			} else {
				i--
			}
		}

		slices.Reverse(path)
		return string(path)
	}

	t := min(f[n-1][n-1][0].x, f[n-1][n-1][0].y, f[n-1][n-1][1].x, f[n-1][n-1][1].y)
	if t <= ans {
		fmt.Fprintln(out, t)
		k := 0
		if min(f[n-1][n-1][0].x, f[n-1][n-1][0].y) > min(f[n-1][n-1][1].x, f[n-1][n-1][1].y) {
			k = 1
		}
		fmt.Fprintln(out, buildPath(n, k))
	} else {
		fmt.Fprintln(out, 1)
		fmt.Fprint(out, strings.Repeat("D", pos[0]))
		fmt.Fprint(out, strings.Repeat("R", pos[1]))
		fmt.Fprint(out, strings.Repeat("D", n-1-pos[0]))
		fmt.Fprintln(out, strings.Repeat("R", n-1-pos[1]))
	}
}
