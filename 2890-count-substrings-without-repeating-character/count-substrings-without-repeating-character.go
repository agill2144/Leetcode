func numberOfSpecialSubstrings(s string) int {
    count := 0
    idx := map[byte]int{}
    left := 0
    for i := 0; i < len(s); i++ {
        val, ok := idx[s[i]]
        if ok && left <= val {
            left = val+1
        }
        idx[s[i]] = i
        count += (i-left+1)
    }
    return count
}