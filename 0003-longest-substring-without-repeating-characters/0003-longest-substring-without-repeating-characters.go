func lengthOfLongestSubstring(s string) int {
    left := 0
    windowState := map[byte]struct{}{}
    windowSize := 0
    for i := 0; i < len(s); i++ {
        char := s[i]
        _, ok := windowState[char]
        for ok {
            leftChar := s[left]
            delete(windowState, leftChar)
            left++
            _, ok = windowState[char]
        }
        windowState[char] = struct{}{}
        if i-left+1 > windowSize {
            windowSize = i-left+1
        }
    }
    return windowSize
}