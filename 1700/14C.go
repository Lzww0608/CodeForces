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

	pos := [4][]int{}
	cnt := make(map[[2]int]bool)
	xMin, xMax, yMin, yMax := math.MaxInt32, math.MinInt32, math.MaxInt32, math.MinInt32
	for i := 0; i < 4; i++ {
		pos[i] = make([]int, 4)
		fmt.Fscan(in, &pos[i][0], &pos[i][1], &pos[i][2], &pos[i][3])
		if pos[i][0] > pos[i][2] || pos[i][0] == pos[i][2] && pos[i][1] > pos[i][3] {
			pos[i][0], pos[i][1], pos[i][2], pos[i][3] = pos[i][2], pos[i][3], pos[i][0], pos[i][1]
		}
		cnt[[2]int{pos[i][0], pos[i][1]}] = true
		cnt[[2]int{pos[i][2], pos[i][3]}] = false
		xMin = min(xMin, pos[i][0], pos[i][2])
		xMax = max(xMax, pos[i][0], pos[i][2])
		yMin = min(yMin, pos[i][1], pos[i][3])
		yMax = max(yMax, pos[i][1], pos[i][3])
	}

	if len(cnt) != 4 {
		fmt.Fprintln(out, "NO")
		return
	}

	ans := [4][]int{
		{xMin, yMin, xMin, yMax},
		{xMin, yMin, xMax, yMin},
		{xMin, yMax, xMax, yMax},
		{xMax, yMin, xMax, yMax},
	}
	a := [4]int{1, 1, 1, 1}
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			f := true
			for k := 0; k < 4; k++ {
				if ans[j][k] != pos[i][k] {
					f = false
					break
				}
			}
			if f {
				a[j]--
				if a[j] < 0 {
					fmt.Fprintln(out, "NO")
					return
				}
			}
		}
	}

	for _, x := range a {
		if x != 0 {
			fmt.Fprintln(out, "NO")
			return
		}
	}

	fmt.Fprintln(out, "YES")

}
