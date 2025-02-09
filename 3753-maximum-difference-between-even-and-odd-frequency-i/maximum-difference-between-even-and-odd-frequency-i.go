func maxDifference(s string) int {
    freq := map[byte]int{}
    for i := 0; i < len(s); i++ {freq[s[i]]++}
    maxO := math.MinInt64
    minE := math.MaxInt64
    for _, v := range freq {
        if v % 2 == 0 {
            minE = min(minE, v)
        } else {
            maxO = max(maxO, v)
        }
    }
    return maxO - minE
}