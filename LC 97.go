func isInterleave(s1 string, s2 string, s3 string) bool {
    n, m := len(s1), len(s2)
    if len(s3) != n + m {
        return false
    }

    f := make([]bool, m + 1)
    f[0] = true
    for i := 0; i <= n; i++ {
        for j := 0; j <= m; j++ {
            k := i + j - 1
            if i > 0 {
                f[j] = f[j] && (s1[i-1] == s3[k])
            }
            if j > 0 {
                f[j] = f[j] || (f[j-1] && (s2[j-1] == s3[k]))
            }
        }
    }

    return f[m]
}