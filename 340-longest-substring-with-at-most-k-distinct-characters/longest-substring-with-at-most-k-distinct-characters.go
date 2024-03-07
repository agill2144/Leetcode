func lengthOfLongestSubstringKDistinct(s string, k int) int {
    maxSize := 0
    left := 0
    freq := map[byte]int{}
    for i := 0; i < len(s); i++ {
        freq[s[i]]++
        for len(freq) > k {
            leftChar := s[left]
            freq[leftChar]--
            if val := freq[leftChar]; val == 0 {
                delete(freq, leftChar)
            }
            left++
        }
        maxSize = max(maxSize, i-left+1)
    }
    return maxSize
}