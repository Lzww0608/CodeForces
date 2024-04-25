func reinitializePermutation(n int) (ans int) {
    vis := make([]bool, n)

    f := func(i int) int {
        if i < n / 2 {
            return i * 2
        }
        return i * 2 - n + 1
    }

    ans = 1
    for i := 0; i < n; i++ {
        vis[i] = true
        cnt, j := 1, f(i)

        for !vis[j] {
            cnt++
            vis[j] = true
            j = f(j)
        }

        ans = lcm(cnt, ans)
    }

    return
}

func lcm(x, y int) int {
    res := x * y

    for y != 0 {
        x, y = y, x % y
    }

    return res / x
}