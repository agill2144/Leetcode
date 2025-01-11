func canConstruct(s string, k int) bool {
    if len(s) == k {return true}
    if len(s) < k {return false}
    set := map[byte]bool{}
    // count 2 size strings
    count := 0
    for i := 0; i < len(s); i++ {
        if set[s[i]] {delete(set, s[i]); count++; continue}
        set[s[i]] = true
    }
    remaining := len(set)
    // our count has 2 size strings with idential chars
    // ex: aa, bb, cc
    // therefore each of those 2 size strings can take 1 more char out of remaining
    // hence the true remaining is whatever is left after taking $count chars out of set
    remaining -= count
    return remaining + count <= k
}