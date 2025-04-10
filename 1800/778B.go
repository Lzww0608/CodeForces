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

	var n, m int
	fmt.Fscan(in, &n, &m)
	mp := make(map[string]int)
	op := make([]byte, n+1)
	A := make([]int, n+1)
	B := make([]int, n+1)
	f := make([]string, n+1)

	for i := 1; i <= n; i++ {
		var s, tmp, t string
		fmt.Fscan(in, &s, &tmp, &t)
		mp[s] = i
		if t[0] == '0' || t[0] == '1' {
			f[i] = t
			op[i] = ' '
		} else {
			var a, b string
			fmt.Fscan(in, &a, &b)
			op[i] = a[0]
			if t[0] == '?' {
				A[i] = n + 1
			} else {
				A[i] = mp[t]
			}

			if b[0] == '?' {
				B[i] = n + 1
			} else {
				B[i] = mp[b]
			}
		}
	}

	//fmt.Fprintln(out, string(op[1:]))
	ansA, ansB := make([]byte, 0, m), make([]byte, 0, m)
	g := make([]int, n+2)
	for i := 0; i < m; i++ {
		sumA, sumB := 0, 0
		g[n+1] = 0
		for j := 1; j <= n; j++ {
			switch op[j] {
			case ' ':
				g[j] = int(f[j][i] - '0')
				break
			case 'X':
				g[j] = g[A[j]] ^ g[B[j]]
				break
			case 'A':
				g[j] = g[A[j]] & g[B[j]]
				break
			case 'O':
				g[j] = g[A[j]] | g[B[j]]
				break
			}
			sumA += g[j]
		}

		g[n+1] = 1
		for j := 1; j <= n; j++ {
			switch op[j] {
			case ' ':

				g[j] = int(f[j][i] - '0')
				break
			case 'X':
				g[j] = g[A[j]] ^ g[B[j]]
				break
			case 'A':
				g[j] = g[A[j]] & g[B[j]]
				break
			case 'O':
				g[j] = g[A[j]] | g[B[j]]
			}
			sumB += g[j]
		}

		if sumA == sumB {
			ansA = append(ansA, '0')
			ansB = append(ansB, '0')
		} else if sumA > sumB {
			ansA = append(ansA, '1')
			ansB = append(ansB, '0')
		} else {
			ansA = append(ansA, '0')
			ansB = append(ansB, '1')
		}
	}
	fmt.Fprintln(out, string(ansA))
	fmt.Fprintln(out, string(ansB))
}
