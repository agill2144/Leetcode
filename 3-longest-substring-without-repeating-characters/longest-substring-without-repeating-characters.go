func lengthOfLongestSubstring(s string) int {
    maxWin := 0
    idx := map[byte]int{}
    left := 0
    for i := 0; i < len(s); i++ {
        lastSeen, ok := idx[s[i]]
        if ok && left <= lastSeen {
            left = lastSeen+1
        }
        idx[s[i]] = i
        maxWin = max(maxWin, i-left+1)
    }
    return maxWin
}