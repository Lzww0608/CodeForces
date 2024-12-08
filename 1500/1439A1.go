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

	var q int
	for fmt.Fscan(in, &q); q > 0; q-- {
		solve(in, out)
	}
}

func solve(in *bufio.Reader, out *bufio.Writer) {
	var n, m int
	fmt.Fscan(in, &n, &m)
	var s string
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int, m)
		fmt.Fscan(in, &s)
		for j := 0; j < m; j++ {
			a[i][j] = int(s[j] - '0')
		}
	}

	ans := [][]int{}
	if n&1 == 1 {
		for j := 0; j < m; j += 2 {
			if j+1 == m {
				if a[n-1][j] == 1 {
					ans = append(ans, []int{n - 1, j, n - 2, j, n - 2, j - 1})
					a[n-1][j] ^= 1
					a[n-2][j] ^= 1
					a[n-2][j-1] ^= 1
				}
			} else {
				if a[n-1][j] == 1 {
					if a[n-1][j+1] == 1 {
						ans = append(ans, []int{n - 1, j, n - 1, j + 1, n - 2, j})
						a[n-1][j] ^= 1
						a[n-1][j+1] ^= 1
						a[n-2][j] ^= 1
					} else {
						ans = append(ans, []int{n - 1, j, n - 2, j, n - 2, j + 1})
						a[n-1][j] ^= 1
						a[n-2][j+1] ^= 1
						a[n-2][j] ^= 1
					}
				} else if a[n-1][j+1] == 1 {
					ans = append(ans, []int{n - 1, j + 1, n - 2, j + 1, n - 2, j})
					a[n-1][j] ^= 1
					a[n-2][j+1] ^= 1
					a[n-2][j] ^= 1
				}
			}
		}
	}

	if m&1 == 1 {
		for i := 0; i < n; i += 2 {
			if i+1 == n {
				if a[i][m-1] == 1 {
					ans = append(ans, []int{i, m - 1, i, m - 2, i - 1, m - 2})
					a[i][m-1] ^= 1
					a[i][m-2] ^= 1
					a[i-1][m-2] ^= 1
				}
			} else {
				if a[i][m-1] == 1 {
					if a[i+1][m-1] == 1 {
						ans = append(ans, []int{i, m - 1, i + 1, m - 1, i + 1, m - 2})
						a[i][m-1] ^= 1
						a[i+1][m-1] ^= 1
						a[i+1][m-2] ^= 1
					} else {
						ans = append(ans, []int{i, m - 1, i, m - 2, i + 1, m - 2})
						a[i][m-1] ^= 1
						a[i][m-2] ^= 1
						a[i+1][m-2] ^= 1
					}
				} else if a[i+1][m-1] == 1 {
					ans = append(ans, []int{i + 1, m - 1, i, m - 2, i + 1, m - 2})
					a[i+1][m-1] ^= 1
					a[i][m-2] ^= 1
					a[i+1][m-2] ^= 1
				}
			}
		}
	}

	change := func(i, j, c, d, e, f int) {
		ans = append(ans, []int{i, j, c, d, e, f})
		a[i][j] ^= 1
		a[c][d] ^= 1
		a[e][f] ^= 1
	}

	for i := 0; i < n-1; i += 2 {
		for j := 0; j < m-1; j += 2 {
			cnt := a[i][j] + a[i+1][j] + a[i][j+1] + a[i+1][j+1]
			if cnt == 0 {
				continue
			} else if cnt == 1 {
				if a[i][j] == 1 {
					change(i, j, i+1, j, i+1, j+1)
					change(i, j, i+1, j, i, j+1)
					change(i, j, i, j+1, i+1, j+1)
				} else if a[i+1][j] == 1 {
					change(i+1, j, i, j, i+1, j+1)
					change(i+1, j, i, j, i, j+1)
					change(i+1, j, i, j+1, i+1, j+1)
				} else if a[i][j+1] == 1 {
					change(i, j+1, i+1, j, i+1, j+1)
					change(i, j+1, i+1, j, i, j)
					change(i, j+1, i, j, i+1, j+1)
				} else {
					change(i+1, j+1, i+1, j, i, j)
					change(i+1, j+1, i+1, j, i, j+1)
					change(i+1, j+1, i, j+1, i, j)
				}
			} else if cnt == 2 {
				if a[i][j]+a[i+1][j] == 2 {
					change(i, j, i, j+1, i+1, j+1)
					change(i+1, j, i, j+1, i+1, j+1)
				} else if a[i][j]+a[i][j+1] == 2 {
					change(i, j, i+1, j, i+1, j+1)
					change(i, j+1, i+1, j, i+1, j+1)
				} else if a[i][j]+a[i+1][j+1] == 2 {
					change(i, j, i+1, j, i, j+1)
					change(i+1, j+1, i+1, j, i, j+1)
				} else if a[i+1][j]+a[i][j+1] == 2 {
					change(i+1, j, i, j, i+1, j+1)
					change(i, j+1, i, j, i+1, j+1)
				} else if a[i+1][j]+a[i+1][j+1] == 2 {
					change(i+1, j, i, j, i, j+1)
					change(i+1, j+1, i, j, i, j+1)
				} else if a[i][j+1]+a[i+1][j+1] == 2 {
					change(i, j+1, i, j, i+1, j)
					change(i+1, j+1, i, j, i+1, j)
				}

			} else if cnt == 3 {
				if a[i][j] == 0 {
					change(i+1, j, i, j+1, i+1, j+1)
				} else if a[i+1][j] == 0 {
					change(i, j, i, j+1, i+1, j+1)
				} else if a[i][j+1] == 0 {
					change(i, j, i+1, j, i+1, j+1)
				} else {
					change(i, j, i+1, j, i, j+1)
				}
			} else if cnt == 4 {
				change(i, j, i+1, j, i, j+1)
				change(i+1, j+1, i+1, j, i, j)
				change(i+1, j+1, i+1, j, i, j+1)
				change(i+1, j+1, i, j+1, i, j)
			}
		}
	}

	fmt.Fprintln(out, len(ans))
	for _, v := range ans {
		for _, x := range v {
			fmt.Fprint(out, x+1, " ")
		}
		fmt.Fprintln(out, "")
	}
}
