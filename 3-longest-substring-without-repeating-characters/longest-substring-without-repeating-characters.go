func lengthOfLongestSubstring(s string) int {
    freq := map[byte]int{}
    maxWin := 0
    left := 0
    for i := 0; i < len(s); i++ {
        freq[s[i]]++
        for freq[s[i]] > 1 {
            freq[s[left]]--
            left++
        }
        maxWin = max(maxWin, i-left+1)
    }
    return maxWin
}