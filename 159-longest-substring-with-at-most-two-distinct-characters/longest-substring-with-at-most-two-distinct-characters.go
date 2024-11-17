
func lengthOfLongestSubstringTwoDistinct(s string) int {
    n := len(s)
    freq := map[byte]int{}
    left := 0
    maxSize := 0
    
    for i := 0; i < n; i++ {
        char := s[i]
        freq[char]++
        for len(freq) > 2 {
            leftChar := s[left]
            freq[leftChar]--
            if freq[leftChar] == 0 {delete(freq, leftChar)}
            left++
        }
        maxSize = max(maxSize, i-left+1)
    }
    return maxSize
}