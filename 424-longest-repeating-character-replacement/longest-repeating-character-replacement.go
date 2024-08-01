func characterReplacement(s string, k int) int {
    maxRepeating := 0
    left := 0
    maxWin := 0
    freq := map[byte]int{}
    for i := 0; i < len(s); i++ {
        freq[s[i]]++
        maxRepeating = max(maxRepeating, freq[s[i]])
        nonRepeating := (i-left+1)-maxRepeating
        if nonRepeating > k {
            freq[s[left]]--
            if freq[s[left]] == 0 {delete(freq, s[left])}
            left++
        } else {
            maxWin = max(maxWin, i-left+1)
        }
    }
    return maxWin
}