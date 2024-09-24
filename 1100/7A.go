package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const N int = 8

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	s := make([]string, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &s[i])
	}

	ans := 0
	for i := 0; i < N; i++ {
		if s[i] == strings.Repeat("B", N) {
			ans++
		}
	}
	if ans == N {
		fmt.Fprintln(out, ans)
		return
	}

	for i := 0; i < N; i++ {
		cnt := 0
		for j := 0; j < N; j++ {
			if s[j][i] == 'B' {
				cnt++
			}
		}
		if cnt == N {
			ans++
		}
	}
	fmt.Fprintln(out, ans)
}
