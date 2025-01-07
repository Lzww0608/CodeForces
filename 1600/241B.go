package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, x int
	fmt.Fscan(in, &n)
	cnt := make([]int, 10)
	sum := 0
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &x)
		cnt[x]++
		sum += x
	}

	if cnt[0] == 0 {
		fmt.Fprintln(out, -1)
		return
	}

	print := func() {
		for i := 1; i <= 9; i++ {
			if cnt[i] > 0 {
				for i := 9; i >= 0; i-- {
					fmt.Fprint(out, strings.Repeat(strconv.Itoa(i), cnt[i]))
				}
				return
			}
		}
		fmt.Fprint(out, 0)
	}

	if sum%3 == 0 {
		print()
	} else {
		a, b := []int{1, 4, 7}, []int{2, 5, 8}
		if sum%3 == 1 {
			for _, i := range a {
				if cnt[i] > 0 {
					cnt[i]--
					print()
					return
				}
			}
			cur := 0
			for _, i := range b {
				cur += cnt[i]
				if cnt[i] > 1 {
					cnt[i] -= 2
					print()
					return
				}
			}
			if cur >= 2 {
				cur = 2
				for i := 0; i <= 2 && cur > 0; i++ {
					if cnt[b[i]] > 0 {
						cnt[b[i]]--
						cur--
					}
				}
				//fmt.Fprintln(out, cnt)
				print()
				return
			}

		} else {
			for _, i := range b {
				if cnt[i] > 0 {
					cnt[i] -= 1
					print()
					return
				}
			}
			cur := 0
			for _, i := range a {
				cur += cnt[i]
				if cnt[i] > 1 {
					cnt[i] -= 2
					print()
					return
				}
			}
			if cur >= 2 {
				cur = 2
				for i := 0; i <= 2 && cur > 0; i++ {
					if cnt[a[i]] > 0 {
						cnt[a[i]]--
						cur--
					}
				}
				print()
				return
			}
		}
		for i := 1; i <= 9; i++ {
			if i%3 != 0 {
				cnt[i] = 0
			}
		}
		print()
		return
	}

}
