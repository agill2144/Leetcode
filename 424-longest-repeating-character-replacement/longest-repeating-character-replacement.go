func characterReplacement(s string, k int) int {
    maxRepeating := 0
    left := 0
    maxWin := 0
    freq := map[byte]int{}
    for i := 0; i < len(s); i++ {
        freq[s[i]]++
        maxRepeating = max(maxRepeating, freq[s[i]])
        nonRepeating := (i-left+1)-maxRepeating
        for nonRepeating > k {
            freq[s[left]]--
            left++
            maxRepeating := 0
            for _, v := range freq {maxRepeating = max(maxRepeating, v)}
            nonRepeating = (i-left+1)-maxRepeating
        }
        maxWin = max(maxWin, i-left+1)
    }
    return maxWin
}