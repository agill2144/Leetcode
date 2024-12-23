func mergeAlternately(word1 string, word2 string) string {
    out := new(strings.Builder)
    m := len(word1)
    n := len(word2)
    for i := 0; i < max(m, n); i++ {
        if i < m {out.WriteByte(word1[i])}
        if i < n {out.WriteByte(word2[i])}
    }
    return out.String()
}

// func mergeAlternately(word1 string, word2 string) string {
//     p1, p2 := 0, 0
//     out := new(strings.Builder)
//     for p1 < len(word1) || p2 < len(word2) {
//         p1Val := ""
//         p2Val := ""
//         if p1 < len(word1) {p1Val = string(word1[p1]); p1++}
//         if p2 < len(word2) {p2Val = string(word2[p2]); p2++}
//         if p1Val != "" {
//             out.WriteString(p1Val)
//         }
//         if p2Val != "" {
//             out.WriteString(p2Val)
//         }
//     }
//     return out.String()
// }