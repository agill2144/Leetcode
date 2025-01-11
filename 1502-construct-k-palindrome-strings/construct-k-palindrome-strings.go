func canConstruct(s string, k int) bool {
    if len(s) == k {return true}
    if len(s) < k {return false}
    set := map[byte]bool{}
    count := 0
    for i := 0; i < len(s); i++ {
        if set[s[i]] {delete(set, s[i]); count++; continue}
        set[s[i]] = true
    }
    left := len(set)
    left -= count
    return left + count <= k    
}