func findThePrefixCommonArray(A []int, B []int) []int {
    aIdx := map[int]int{}
    bIdx := map[int]int{}
    n := len(A)
    for i := 0; i < n; i++ {
        aIdx[A[i]] = i
        bIdx[B[i]] = i
    }
    cp := make([]int, n)
    for i := 0; i < n; i++ {
        if i-1 >= 0 {cp[i] = cp[i-1]}
        a := A[i]
        b := B[i]
        if bIdx[a] <= i {
            cp[i]++
        }
        if a == b {continue}
        if aIdx[b] <= i {cp[i]++}
    }
    return cp
}

