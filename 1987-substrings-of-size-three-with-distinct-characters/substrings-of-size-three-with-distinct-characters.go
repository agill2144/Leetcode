func countGoodSubstrings(s string) int {
    freq := map[byte]int{}
    count := 0
    left := 0
    for i := 0; i < len(s); i++ {
        freq[s[i]]++
        for freq[s[i]] > 1 {
            if val, ok := freq[s[left]]; ok {
                freq[s[left]]--
                if val == 1 {delete(freq, s[left])}
            }
            left++
        }
        if i-left+1 == 3 {
            count++
            if val, ok := freq[s[left]]; ok {
                freq[s[left]]--
                if val == 1 {delete(freq, s[left])}
            }
            left++
        }
    }
    return count
}