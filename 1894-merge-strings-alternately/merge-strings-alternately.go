func mergeAlternately(word1 string, word2 string) string {
    p1, p2 := 0, 0
    out := new(strings.Builder)
    for p1 < len(word1) || p2 < len(word2) {
        p1Val := ""
        p2Val := ""
        if p1 < len(word1) {p1Val = string(word1[p1]); p1++}
        if p2 < len(word2) {p2Val = string(word2[p2]); p2++}
        if p1Val != "" {
            out.WriteString(p1Val)
        }
        if p2Val != "" {
            out.WriteString(p2Val)
        }
    }
    return out.String()
}