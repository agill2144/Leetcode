func characterReplacement(s string, k int) int {
    maxWin := 0
    left := 0
    maxRepeating := 0
    freq := map[byte]int{}
    for i := 0; i < len(s); i++ {
        freq[s[i]]++
        maxRepeating = max(maxRepeating, freq[s[i]])
        size := i-left+1
        nonRepeating := size - maxRepeating
        if nonRepeating <= k {
            maxWin = max(maxWin, size)
        } else {
            freq[s[left]]--
            left++
        }
    }
    return maxWin
}