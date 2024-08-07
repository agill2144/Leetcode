func numberOfSpecialSubstrings(s string) int {
    idx := map[byte]int{}
    count := 0
    left := 0
    for i := 0; i < len(s); i++ {
        lastSeenAt, ok := idx[s[i]]
        if ok {
            if left <= lastSeenAt {
                left = lastSeenAt+1
            }
        }
        idx[s[i]] = i
        count += (i-left+1)
    }
    return count
}