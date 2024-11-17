
func lengthOfLongestSubstringTwoDistinct(s string) int {
    n := len(s)
    freq := map[byte]int{}
    left := 0
    maxSize := 0
    
    for i := 0; i < n; i++ {
        char := s[i]
        freq[char]++
        if len(freq) <= 2 {
            maxSize = max(maxSize, i-left+1)
        } else {
            freq[s[left]]--
            if freq[s[left]] == 0 {delete(freq, s[left])}
            left++
        }
        
    }
    return maxSize
}