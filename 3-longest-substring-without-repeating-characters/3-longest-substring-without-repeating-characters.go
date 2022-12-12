func lengthOfLongestSubstring(s string) int {
    if len(s) <= 1 {return len(s)}
    windowState := map[byte]int{}
    left := 0
    max := 0
    for right := 0; right < len(s); right++ {
        rightChar := s[right]
        lastSeenAtIdx, ok := windowState[rightChar]
        if ok {
            if left <= lastSeenAtIdx {
                left = lastSeenAtIdx+1
            }
        }
        windowState[rightChar] = right
        if right-left+1 > max {
            max = right-left+1
        }
    }
    return max
}