func lengthOfLongestSubstringTwoDistinct(s string) int {
    left := 0
    freq := map[byte]int{}
    maxWin := 0
    for i := 0; i < len(s); i++ {
        freq[s[i]]++
        if len(freq) <= 2 {
            maxWin = max(maxWin, i-left+1)
        } else {
            freq[s[left]]--
            if freq[s[left]] == 0{delete(freq, s[left])}
            left++
        }
    }
    return maxWin
}