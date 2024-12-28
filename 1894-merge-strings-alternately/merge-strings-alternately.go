func mergeAlternately(word1 string, word2 string) string {
    out := new(strings.Builder)
    w1, w2 := 0,0 
    for w1 < len(word1) && w2 < len(word2) {
        out.WriteByte(word1[w1])
        w1++
        out.WriteByte(word2[w2])
        w2++
    }
    for w1 < len(word1) {
        out.WriteByte(word1[w1])
        w1++
    }
    for w2 < len(word2) {
        out.WriteByte(word2[w2])
        w2++
    }

    return out.String()
}