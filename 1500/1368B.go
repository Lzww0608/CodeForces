package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var k int
	fmt.Fscan(in, &k)
	if k == 1 {
		fmt.Fprintln(out, "codeforces")
		return
	}
	m := []byte{
		'c',
		'o',
		'd',
		'e',
		'f',
		'o',
		'r',
		'c',
		'e',
		's',
	}
	cnt, sum := 1, 1
	for sum < k {
		cnt++
		sum = quickPow(cnt, 10)
	}

	t := 9
	for t > 0 && quickPow(cnt, t)*quickPow(cnt-1, 10-t) >= k {
		t--
	}
	t++
	for i := 0; i < t; i++ {
		fmt.Fprint(out, strings.Repeat(string(m[i]), cnt))
	}
	for i := t; i < 10; i++ {
		fmt.Fprint(out, strings.Repeat(string(m[i]), cnt-1))
	}

}

func quickPow(a, r int) int {
	res := 1
	for r > 0 {
		if r&1 == 1 {
			res *= a
		}
		a *= a
		r >>= 1
	}

	return res
}
