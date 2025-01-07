func mergeAlternately(word1 string, word2 string) string {
    res := new(strings.Builder)
    for i := 0; i < max(len(word1),len(word2)); i++ {
        if i < len(word1) {
            res.WriteByte(word1[i])
        }
        if i < len(word2) {
            res.WriteByte(word2[i])
        }
    }
    return res.String()
}