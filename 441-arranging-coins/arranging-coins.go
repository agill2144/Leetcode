func arrangeCoins(n int) int {
    k := 1
    for n >= k {
        if n < k {break}
        n -= k
        k++
    }
    return k-1
}