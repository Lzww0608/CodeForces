func maximumGap(nums []int) (ans int) {
    type pair struct {
        mx, mn int 
    }
    n := len(nums)
    if n < 2 {
        return 0
    }
    mn := slices.Min(nums)
    mx := slices.Max(nums)
    d := max(1, (mx - mn) / (n - 1))
    buckets := make([]pair, (mx - mn) / d + 1)
    for i := range buckets {
        buckets[i] = pair{-1, -1}
    }

    for _, x := range nums {
        pos := (x - mn) / d 
        if buckets[pos].mx == -1 {
            buckets[pos] = pair{x, x}
        } else {
            buckets[pos].mn = min(buckets[pos].mn, x)
            buckets[pos].mx = max(buckets[pos].mx, x)
        }
    }

    pre := -1
    for _, v := range buckets {
        if v.mn == -1 {
            continue
        }
        if pre != -1 {
            ans = max(ans, v.mn - pre)
        }
        pre = v.mx
    }

    return
}