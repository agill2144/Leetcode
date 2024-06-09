func valueAfterKSeconds(n int, k int) int {
    if k <= 0 {return 1}
    mod := 1000000007
    tmp := make([]int, n)
    for i := 0; i < n; i++ {
        tmp[i] = 1
    }

    for k != 0 {
        for i := 1; i < n; i++ {
            curr := tmp[i]
            prev := tmp[i-1]
            tmp[i] = (curr+prev) % mod
        }
        k--
    }
    return tmp[n-1]
}