func characterReplacement(s string, k int) int {
    maxWin := 0
    left := 0
    freq := map[byte]int{}
    maxFreq := 0
    for i := 0; i < len(s); i++ {
        freq[s[i]]++
        maxFreq = max(maxFreq, freq[s[i]])
        if i-left+1 - maxFreq <= k {
            maxWin = max(maxWin, i-left+1)
        } else {
            freq[s[left]]--
            if freq[s[left]] == 0 {delete(freq, s[left])}
            left++
        }
    }
    return maxWin
}