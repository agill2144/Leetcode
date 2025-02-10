func lengthOfLongestSubstringKDistinct(s string, k int) int {
    freq := map[byte]int{}
    left := 0
    maxWin := 0
    for i := 0; i < len(s); i++ {
        freq[s[i]]++
        if len(freq) <= k {
            maxWin = max(maxWin, i-left+1)
        } else {
            freq[s[left]]--
            if freq[s[left]] == 0 {
                delete(freq, s[left])
            }
            left++
        }    
    }
    return maxWin
}