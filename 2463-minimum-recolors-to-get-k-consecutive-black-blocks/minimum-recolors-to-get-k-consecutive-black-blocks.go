func minimumRecolors(blocks string, k int) int {
    n := len(blocks)
    left := 0
    bCount := 0
    res := math.MaxInt64
    for i := 0; i < n; i++ {
        if blocks[i] == 'B' {bCount++}
        if i-left+1 == k {
            res = min(res, (i-left+1)-bCount)
            if blocks[left] == 'B' {bCount--}
            left++
        }
    }
    return res
}