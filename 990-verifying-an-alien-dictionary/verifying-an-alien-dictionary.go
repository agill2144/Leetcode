func isAlienSorted(words []string, order string) bool {
    /*
        abcd
        ab
    */
    idx := map[byte]int{}
    for i := 0; i < len(order); i++ {idx[order[i]]=i}
    for i := 1; i < len(words); i++ {
        prev := words[i-1]
        curr := words[i]
        c, p := 0,0
        for c < len(curr) && p < len(prev) {
            cChar := curr[c]
            pChar := prev[p]
            if cChar == pChar {
                c++; p++
            } else {
                if idx[pChar] > idx[cChar] {return false}
                break
            }
        }
        if c == len(curr) && p < len(prev) {return false}
    }
    return true
}