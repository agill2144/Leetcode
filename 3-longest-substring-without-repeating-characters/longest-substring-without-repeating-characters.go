func lengthOfLongestSubstring(s string) int {
    res := 0
    left := 0
    freq := map[byte]int{}
    for i := 0; i < len(s); i++ {
        freq[s[i]]++
        for freq[s[i]] > 1 {
            freq[s[left]]--
            if freq[s[left]] == 0 {delete(freq, s[left])}
            left++
        }
        res = max(res, i-left+1)
    }
    return res
}