func longestPalindromeSubseq(s string) int {
    n := len(s)
    f := make([]int, n)

    for i := n - 1; i >= 0; i-- {
        pre := f[i]
        f[i] = 1
        for j := i + 1; j < n; j++ {
            if s[i] == s[j] {
                pre, f[j] = f[j], pre + 2
            } else {
                pre = f[j]
                f[j] = max(f[j], f[j-1])
            }
        }
    }

    return f[n-1]
}