func lengthOfLongestSubstring(s string) int {
    res := 0
    left := 0
    idx := map[byte]int{}
    for i := 0; i < len(s); i++ {
        val, ok := idx[s[i]]
        if ok && left <= val {
            left = val+1
        }
        idx[s[i]] = i
        res = max(res, i-left+1)
    }
    return res
}