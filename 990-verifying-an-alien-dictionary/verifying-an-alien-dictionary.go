func isAlienSorted(words []string, order string) bool {
    idx := map[byte]int{}
    for i := 0; i < len(order); i++ {idx[order[i]] = i}
    for i := 1; i < len(words); i++ {
        curr, prev := words[i], words[i-1]
        c, p := 0,0
        for c < len(curr) && p < len(prev) {
            if curr[c] == prev[c] {c++; p++; continue}
            cIdx := idx[curr[c]]
            pIdx := idx[prev[p]]
            if pIdx > cIdx {return false}
            break
        }

        if p < len(prev) && c == len(curr) {return false}
    }
    return true
}