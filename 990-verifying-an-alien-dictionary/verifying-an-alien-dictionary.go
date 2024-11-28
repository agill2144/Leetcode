// we need to check whether words are sorted accoriding to a custom character order
func isAlienSorted(words []string, order string) bool {
    if len(words) <= 1 {return true}
    idx := map[byte]int{}
    for i := 0; i < len(order); i++ {idx[order[i]] = i}
    for i := 1; i < len(words); i++ {
        curr := words[i]
        prev := words[i-1]
        c, p := 0,0
        for c < len(curr) && p < len(prev) {
            cChar := curr[c]
            pChar := prev[p]
            if cChar == pChar {p++; c++; continue}
            cCharIdx := idx[cChar]
            pCharIdx := idx[pChar]
            if cCharIdx < pCharIdx {return false} 
            break
        }
        if p < len(prev) && c == len(curr) {return false}
    }
    return true
}
