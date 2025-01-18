package main

import (
	"bufio"
	"fmt"
	"os"
)

const N int = 5_000_001

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	minPrime := make([]int, N+1) // 存储每个数的最小质因子
	primes := make([]int, 0)     // 存储所有素数
	for i := 2; i <= N; i++ {
		if minPrime[i] == 0 { // i 是素数
			minPrime[i] = i
			primes = append(primes, i)
		}
		// 遍历已知素数
		for _, p := range primes {
			if p > minPrime[i] || i*p > N {
				break
			}
			minPrime[i*p] = p // i * p 的最小质因子是 p
		}
	}

	// 计算每个数的质因数指数之和
	factor := make([]int, N+1) // 质因数指数之和数组
	for i := 2; i <= N; i++ {
		if minPrime[i] == i { // i 是素数
			factor[i] = 1 // 素数的指数之和为 1
		} else {
			// 对 i 进行质因数分解
			x := i
			p := minPrime[x]
			for x%p == 0 {
				x /= p
				factor[i]++
			}
			// 加上剩余部分的指数之和
			factor[i] += factor[x]
		}
	}

	for i := 1; i < N; i++ {
		factor[i] += factor[i-1]
	}

	var t, a, b int
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &a, &b)
		fmt.Fprintln(out, factor[a]-factor[b])
		//fmt.Fprintln(out, factor[a], factor[b])
	}
}
