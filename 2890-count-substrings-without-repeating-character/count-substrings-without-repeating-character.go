func numberOfSpecialSubstrings(s string) int {
    freq := map[byte]int{}
    count := 0
    left := 0
    for i := 0; i < len(s); i++ {
        _, ok := freq[s[i]]
        for ok {
            leftChar := s[left]
            freq[leftChar]--
            if freq[leftChar] == 0 {delete(freq, leftChar)}
            left++
            _, ok = freq[s[i]]
        }
        freq[s[i]]++
        count += (i-left+1)
    }
    return count
}