func lengthOfLongestSubstringKDistinct(s string, k int) int {
    freq := map[byte]int{}
    maxSize := 0
    left := 0
    for i := 0; i < len(s); i++ {
        freq[s[i]]++
        if len(freq) <= k {
            maxSize = max(maxSize, i-left+1)
        } else {
            freq[s[left]]--
            if freq[s[left]] == 0 {delete(freq, s[left])}
            left++
        }
    }
    return maxSize
}