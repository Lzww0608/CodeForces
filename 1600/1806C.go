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

	var t int
	for fmt.Fscan(in, &t); t > 0; t-- {
		solve(in, out)
	}

}
func solve(in *bufio.Reader, out *bufio.Writer) {
	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n<<1)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}
	if n == 1 {
		fmt.Fprintln(out, abs(a[1]-a[0]))
		return
	} else if n == 2 {
		ans := min(abs(a[0])+abs(a[1])+abs(a[2])+abs(a[3]), abs(a[0]-2)+abs(a[1]-2)+abs(a[2]-2)+abs(a[3]-2))
		for i := 0; i < n*2; i++ {
			cur := 0
			for j := 0; j < n*2; j++ {
				if i == j {
					cur += abs(a[j] - 2)
				} else {
					cur += abs(a[j] + 1)
				}
			}
			ans = min(ans, cur)
		}
		fmt.Fprintln(out, ans)
		return
	}

	if n&1 == 1 {
		ans := 0
		for _, x := range a {
			ans += abs(x)
		}
		fmt.Fprintln(out, ans)
		return
	} else {
		ans := 0
		for _, x := range a {
			ans += abs(x)
		}
		cur := 0
		for i := 0; i < n*2; i++ {
			cur += abs(a[i] + 1)
		}
		for i := 0; i < n*2; i++ {
			ans = min(ans, cur+abs(a[i]-n)-abs(a[i]+1))
		}
		fmt.Fprintln(out, ans)
		return
	}

}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
