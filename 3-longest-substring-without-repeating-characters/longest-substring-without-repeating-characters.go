func lengthOfLongestSubstring(s string) int {
    idx := map[byte]int{}
    left := 0
    maxWin := 0
    for i := 0; i < len(s); i++ {
        val, ok := idx[s[i]]
        if ok {
            if left <= val {
                left = val+1
            }
        }
        maxWin = max(maxWin, i-left+1)
        idx[s[i]] = i
    }
    return maxWin
}