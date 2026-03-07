func countSubarrays(nums []int, k int, m int) int64 {
    n := len(nums)
    
    solve := func(t int) (res int) {
        cnt := make(map[int]int)
        s := 0
        for l, r := 0, 0; r < n; r++ {
            x := nums[r]
            if cnt[x]++; cnt[x] == m {
                s++
            }

            for len(cnt) >= t && s >= k {
                y := nums[l]
                cnt[y]--
                if cnt[y] + 1 == m {
                    s--
                }
                if cnt[y] == 0 {
                    delete(cnt, y)
                }
                l++
            }
            res += l
        }

        return
    }

    return int64(solve(k) - solve(k + 1))
}