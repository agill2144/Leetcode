func lengthOfLongestSubstring(s string) int {
    left := 0
    idx := map[byte]int{}
    maxWin := 0
    for i := 0; i < len(s); i++ {
        lastIdx, ok := idx[s[i]]
        if ok {
            if left <= lastIdx {left = lastIdx+1}
        }
        idx[s[i]] = i
        maxWin = max(maxWin, i-left+1)
    }
    return maxWin
}