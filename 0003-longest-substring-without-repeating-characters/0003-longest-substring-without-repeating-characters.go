func lengthOfLongestSubstring(s string) int {
    left := 0
    maxSize := 0
    charIdxMap := map[byte]int{}
    for i := 0; i < len(s); i++ {
        char := s[i]
        idxLastSeenAt, ok := charIdxMap[char]
        if ok {
            if left <= idxLastSeenAt {
                left = idxLastSeenAt+1
            }
        }
        charIdxMap[s[i]] = i
        if i-left+1 > maxSize {
            maxSize = i-left+1
        }
    }
    return maxSize
}