func lengthOfLongestSubstringKDistinct(s string, k int) int {
    // we are allowed at have <= k uniq chars in a window
    freq := map[byte]int{}
    left := 0
    maxWin := 0
    for i := 0; i < len(s); i++ {
        freq[s[i]]++
        for len(freq) > k {
            freq[s[left]]--
            if freq[s[left]] == 0 {delete(freq, s[left])}
            left++
        }
        maxWin = max(maxWin, i-left+1)
    }
    return maxWin
}