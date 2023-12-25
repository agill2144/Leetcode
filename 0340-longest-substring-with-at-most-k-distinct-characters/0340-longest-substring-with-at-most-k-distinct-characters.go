func lengthOfLongestSubstringKDistinct(s string, k int) int {
    left := 0
    windowMap := map[byte]int{}
    maxLen := 0
    for i := 0; i < len(s); i++ {
        char := s[i]
        windowMap[char]++
        for len(windowMap) > k {
            leftChar := s[left]
            windowMap[leftChar]--
            if val := windowMap[leftChar]; val == 0 {
                delete(windowMap, leftChar)
            }
            left++
        }
        if i-left+1 > maxLen {
            maxLen = i-left+1
        }
    }
    return maxLen
}