func numberOfSubstrings(s string) int {
    freq := map[byte]int{}
    left := 0
    count := 0
    n := len(s)
    for i := 0; i < n; i++ {
        freq[s[i]]++
        for freq['a'] > 0 && freq['b'] > 0 && freq['c'] > 0 {
            count += ((n-1)-i+1)
            freq[s[left]]--
            left++
        }
    }
    return count
}